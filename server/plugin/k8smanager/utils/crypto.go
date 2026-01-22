package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

var (
	// ErrInvalidCiphertext 密文格式错误
	ErrInvalidCiphertext = errors.New("invalid ciphertext")
	// ErrDecryptionFailed 解密失败
	ErrDecryptionFailed = errors.New("decryption failed")
)

// getEncryptionKey 获取加密密钥
// 优先从环境变量获取，如果没有则使用配置文件中的密钥
func getEncryptionKey() []byte {
	// 从环境变量获取密钥
	key := global.GVA_CONFIG.K8sManager.EncryptionKey
	if key == "" {
		// 如果没有配置密钥，使用默认密钥（仅用于开发环境，生产环境必须配置）
		global.GVA_LOG.Warn("K8sManager encryption key not configured, using default key (not recommended for production)")
		key = "tianqi-gpu-manage-default-key-32bytes!!"
	}

	// 确保32字节密钥（AES-256）
	keyBytes := []byte(key)
	if len(keyBytes) < 32 {
		// 填充到32字节
		padded := make([]byte, 32)
		copy(padded, keyBytes)
		keyBytes = padded
	}
	return keyBytes[:32]
}

// Encrypt 使用 AES-GCM 加密数据
func Encrypt(plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil
	}

	key := getEncryptionKey()

	// 创建 AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		global.GVA_LOG.Error("Failed to create cipher", zap.Error(err))
		return "", err
	}

	// 创建 GCM 模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		global.GVA_LOG.Error("Failed to create GCM", zap.Error(err))
		return "", err
	}

	// 生成随机 nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		global.GVA_LOG.Error("Failed to generate nonce", zap.Error(err))
		return "", err
	}

	// 加密数据
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Base64 编码
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 使用 AES-GCM 解密数据
func Decrypt(ciphertext string) (string, error) {
	if ciphertext == "" {
		return "", nil
	}

	key := getEncryptionKey()

	// Base64 解码
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		global.GVA_LOG.Error("Failed to decode base64", zap.Error(err))
		return "", ErrInvalidCiphertext
	}

	// 创建 AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		global.GVA_LOG.Error("Failed to create cipher", zap.Error(err))
		return "", err
	}

	// 创建 GCM 模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		global.GVA_LOG.Error("Failed to create GCM", zap.Error(err))
		return "", err
	}

	// 检查长度
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", ErrInvalidCiphertext
	}

	// 提取 nonce 和密文
	nonce, cipherData := data[:nonceSize], data[nonceSize:]

	// 解密
	plaintext, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		global.GVA_LOG.Error("Failed to decrypt", zap.Error(err))
		return "", ErrDecryptionFailed
	}

	return string(plaintext), nil
}

// MustEncrypt 加密数据，如果失败则 panic（用于初始化阶段）
func MustEncrypt(plaintext string) string {
	result, err := Encrypt(plaintext)
	if err != nil {
		panic(err)
	}
	return result
}

// MustDecrypt 解密数据，如果失败则 panic（用于初始化阶段）
func MustDecrypt(ciphertext string) string {
	result, err := Decrypt(ciphertext)
	if err != nil {
		panic(err)
	}
	return result
}
