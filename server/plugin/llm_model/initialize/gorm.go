package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/llm_model/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func Gorm(ctx context.Context) {
	err := global.GVA_DB.WithContext(ctx).AutoMigrate(model.LlmModel{})
	if err != nil {
		err = errors.Wrap(err, "注册表失败!")
		zap.L().Error(fmt.Sprintf("%+v", err))
	}
	InitData(ctx)
}

func ptr(s string) *string {
	return &s
}

func InitData(ctx context.Context) {
	var count int64
	global.GVA_DB.Model(&model.LlmModel{}).Count(&count)
	if count > 0 {
		return
	}
	data := []model.LlmModel{
		{Name: ptr("Qwen-7B-Chat"), Description: ptr("阿里云通义千问70亿参数对话模型"), Url: ptr("https://modelscope.cn/models/qwen/Qwen-7B-Chat/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("Qwen")},
		{Name: ptr("Qwen-14B-Chat"), Description: ptr("阿里云通义千问140亿参数对话模型"), Url: ptr("https://modelscope.cn/models/qwen/Qwen-14B-Chat/summary"), Type: "general_llm", Parameters: ptr("14B"), Publisher: ptr("Qwen")},
		{Name: ptr("Qwen-72B-Chat"), Description: ptr("阿里云通义千问720亿参数对话模型"), Url: ptr("https://modelscope.cn/models/qwen/Qwen-72B-Chat/summary"), Type: "general_llm", Parameters: ptr("72B"), Publisher: ptr("Qwen")},
		{Name: ptr("Qwen1.5-7B-Chat"), Description: ptr("阿里云通义千问1.5版本70亿参数对话模型"), Url: ptr("https://modelscope.cn/models/qwen/Qwen1.5-7B-Chat/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("Qwen")},
		{Name: ptr("Qwen1.5-14B-Chat"), Description: ptr("阿里云通义千问1.5版本140亿参数对话模型"), Url: ptr("https://modelscope.cn/models/qwen/Qwen1.5-14B-Chat/summary"), Type: "general_llm", Parameters: ptr("14B"), Publisher: ptr("Qwen")},
		{Name: ptr("Qwen1.5-32B-Chat"), Description: ptr("阿里云通义千问1.5版本320亿参数对话模型"), Url: ptr("https://modelscope.cn/models/qwen/Qwen1.5-32B-Chat/summary"), Type: "general_llm", Parameters: ptr("32B"), Publisher: ptr("Qwen")},
		{Name: ptr("Qwen1.5-72B-Chat"), Description: ptr("阿里云通义千问1.5版本720亿参数对话模型"), Url: ptr("https://modelscope.cn/models/qwen/Qwen1.5-72B-Chat/summary"), Type: "general_llm", Parameters: ptr("72B"), Publisher: ptr("Qwen")},
		{Name: ptr("Qwen1.5-110B-Chat"), Description: ptr("阿里云通义千问1.5版本1100亿参数对话模型"), Url: ptr("https://modelscope.cn/models/qwen/Qwen1.5-110B-Chat/summary"), Type: "general_llm", Parameters: ptr("110B"), Publisher: ptr("Qwen")},
		{Name: ptr("Baichuan2-7B-Chat"), Description: ptr("百川智能第二代70亿参数对话模型"), Url: ptr("https://modelscope.cn/models/baichuan-inc/Baichuan2-7B-Chat/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("Baichuan")},
		{Name: ptr("Baichuan2-13B-Chat"), Description: ptr("百川智能第二代130亿参数对话模型"), Url: ptr("https://modelscope.cn/models/baichuan-inc/Baichuan2-13B-Chat/summary"), Type: "general_llm", Parameters: ptr("13B"), Publisher: ptr("Baichuan")},
		{Name: ptr("ChatGLM3-6B"), Description: ptr("智谱AI第三代60亿参数对话模型"), Url: ptr("https://modelscope.cn/models/ZhipuAI/chatglm3-6b/summary"), Type: "general_llm", Parameters: ptr("6B"), Publisher: ptr("ZhipuAI")},
		{Name: ptr("GLM-4-9B-Chat"), Description: ptr("智谱AI第四代90亿参数对话模型"), Url: ptr("https://modelscope.cn/models/ZhipuAI/glm-4-9b-chat/summary"), Type: "general_llm", Parameters: ptr("9B"), Publisher: ptr("ZhipuAI")},
		{Name: ptr("Llama-3-8B-Instruct"), Description: ptr("Meta Llama3 80亿参数指令微调模型"), Url: ptr("https://modelscope.cn/models/LLM-Research/Meta-Llama-3-8B-Instruct/summary"), Type: "general_llm", Parameters: ptr("8B"), Publisher: ptr("Meta")},
		{Name: ptr("Llama-3-70B-Instruct"), Description: ptr("Meta Llama3 700亿参数指令微调模型"), Url: ptr("https://modelscope.cn/models/LLM-Research/Meta-Llama-3-70B-Instruct/summary"), Type: "general_llm", Parameters: ptr("70B"), Publisher: ptr("Meta")},
		{Name: ptr("Yi-6B-Chat"), Description: ptr("零一万物60亿参数对话模型"), Url: ptr("https://modelscope.cn/models/01ai/Yi-6B-Chat/summary"), Type: "general_llm", Parameters: ptr("6B"), Publisher: ptr("01.AI")},
		{Name: ptr("Yi-34B-Chat"), Description: ptr("零一万物340亿参数对话模型"), Url: ptr("https://modelscope.cn/models/01ai/Yi-34B-Chat/summary"), Type: "general_llm", Parameters: ptr("34B"), Publisher: ptr("01.AI")},
		{Name: ptr("DeepSeek-V2-Lite-Chat"), Description: ptr("深度求索V2 Lite对话模型"), Url: ptr("https://modelscope.cn/models/deepseek-ai/DeepSeek-V2-Lite-Chat/summary"), Type: "general_llm", Parameters: ptr("16B"), Publisher: ptr("DeepSeek")},
		{Name: ptr("DeepSeek-Coder-V2-Lite-Instruct"), Description: ptr("深度求索代码生成模型V2 Lite"), Url: ptr("https://modelscope.cn/models/deepseek-ai/DeepSeek-Coder-V2-Lite-Instruct/summary"), Type: "general_llm", Parameters: ptr("16B"), Publisher: ptr("DeepSeek")},
		{Name: ptr("InternLM2.5-7B-Chat"), Description: ptr("书生·浦语2.5版本70亿参数对话模型"), Url: ptr("https://modelscope.cn/models/Shanghai_AI_Laboratory/internlm2_5-7b-chat/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("Shanghai AI Lab")},
		{Name: ptr("InternLM2.5-20B-Chat"), Description: ptr("书生·浦语2.5版本200亿参数对话模型"), Url: ptr("https://modelscope.cn/models/Shanghai_AI_Laboratory/internlm2_5-20b-chat/summary"), Type: "general_llm", Parameters: ptr("20B"), Publisher: ptr("Shanghai AI Lab")},
		{Name: ptr("Mistral-7B-Instruct-v0.3"), Description: ptr("Mistral AI 7B 指令微调模型"), Url: ptr("https://modelscope.cn/models/LLM-Research/Mistral-7B-Instruct-v0.3/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("Mistral AI")},
		{Name: ptr("Gemma-2-9b-it"), Description: ptr("Google Gemma2 90亿参数指令微调模型"), Url: ptr("https://modelscope.cn/models/LLM-Research/gemma-2-9b-it/summary"), Type: "general_llm", Parameters: ptr("9B"), Publisher: ptr("Google")},
		{Name: ptr("Gemma-2-27b-it"), Description: ptr("Google Gemma2 270亿参数指令微调模型"), Url: ptr("https://modelscope.cn/models/LLM-Research/gemma-2-27b-it/summary"), Type: "general_llm", Parameters: ptr("27B"), Publisher: ptr("Google")},
		{Name: ptr("Phi-3-mini-4k-instruct"), Description: ptr("Microsoft Phi-3 mini 4k上下文指令微调模型"), Url: ptr("https://modelscope.cn/models/LLM-Research/Phi-3-mini-4k-instruct/summary"), Type: "general_llm", Parameters: ptr("3.8B"), Publisher: ptr("Microsoft")},
		{Name: ptr("Phi-3-medium-4k-instruct"), Description: ptr("Microsoft Phi-3 medium 4k上下文指令微调模型"), Url: ptr("https://modelscope.cn/models/LLM-Research/Phi-3-medium-4k-instruct/summary"), Type: "general_llm", Parameters: ptr("14B"), Publisher: ptr("Microsoft")},
		{Name: ptr("Orion-14B-Chat"), Description: ptr("猎户星空140亿参数对话模型"), Url: ptr("https://modelscope.cn/models/OrionStarAI/Orion-14B-Chat/summary"), Type: "general_llm", Parameters: ptr("14B"), Publisher: ptr("OrionStar")},
		{Name: ptr("XVERSE-13B-Chat"), Description: ptr("元象XVERSE 130亿参数对话模型"), Url: ptr("https://modelscope.cn/models/xverse/XVERSE-13B-Chat/summary"), Type: "general_llm", Parameters: ptr("13B"), Publisher: ptr("XVERSE")},
		{Name: ptr("XVERSE-65B-Chat"), Description: ptr("元象XVERSE 650亿参数对话模型"), Url: ptr("https://modelscope.cn/models/xverse/XVERSE-65B-Chat/summary"), Type: "general_llm", Parameters: ptr("65B"), Publisher: ptr("XVERSE")},
		{Name: ptr("BlueLM-7B-Chat"), Description: ptr("vivo蓝心大模型70亿参数对话模型"), Url: ptr("https://modelscope.cn/models/vivo-ai/BlueLM-7B-Chat/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("vivo")},
		{Name: ptr("MiniCPM-Llama3-V-2_5"), Description: ptr("面壁智能多模态大模型"), Url: ptr("https://modelscope.cn/models/OpenBMB/MiniCPM-Llama3-V-2_5/summary"), Type: "multimodal", Parameters: ptr("8B"), Publisher: ptr("OpenBMB")},
		{Name: ptr("Qwen-VL-Chat"), Description: ptr("阿里云通义千问视觉语言模型"), Url: ptr("https://modelscope.cn/models/qwen/Qwen-VL-Chat/summary"), Type: "vision", Parameters: ptr("7B"), Publisher: ptr("Qwen")},
		{Name: ptr("Yi-VL-34B"), Description: ptr("零一万物视觉语言模型"), Url: ptr("https://modelscope.cn/models/01ai/Yi-VL-34B/summary"), Type: "vision", Parameters: ptr("34B"), Publisher: ptr("01.AI")},
		{Name: ptr("DeepSeek-VL-7B-chat"), Description: ptr("深度求索视觉语言模型"), Url: ptr("https://modelscope.cn/models/deepseek-ai/deepseek-vl-7b-chat/summary"), Type: "vision", Parameters: ptr("7B"), Publisher: ptr("DeepSeek")},
		{Name: ptr("InternVL2-26B"), Description: ptr("书生·浦语视觉语言模型"), Url: ptr("https://modelscope.cn/models/OpenGVLab/InternVL2-26B/summary"), Type: "vision", Parameters: ptr("26B"), Publisher: ptr("OpenGVLab")},
		{Name: ptr("CogVLM2-Llama3-Chat-19B"), Description: ptr("智谱AI视觉语言模型"), Url: ptr("https://modelscope.cn/models/ZhipuAI/CogVLM2-Llama3-Chat-19B/summary"), Type: "vision", Parameters: ptr("19B"), Publisher: ptr("ZhipuAI")},
		{Name: ptr("Qwen-Audio-Chat"), Description: ptr("阿里云通义千问音频模型"), Url: ptr("https://modelscope.cn/models/qwen/Qwen-Audio-Chat/summary"), Type: "audio", Parameters: ptr("7B"), Publisher: ptr("Qwen")},
		{Name: ptr("Whisper-large-v3"), Description: ptr("OpenAI Whisper语音识别模型"), Url: ptr("https://modelscope.cn/models/AI-ModelScope/whisper-large-v3/summary"), Type: "audio", Parameters: ptr("Large"), Publisher: ptr("OpenAI")},
		{Name: ptr("SenseVoiceSmall"), Description: ptr("阿里通义实验室语音识别模型"), Url: ptr("https://modelscope.cn/models/iic/SenseVoiceSmall/summary"), Type: "audio", Parameters: ptr("Small"), Publisher: ptr("Alibaba")},
		{Name: ptr("CosyVoice-300M"), Description: ptr("阿里通义实验室语音合成模型"), Url: ptr("https://modelscope.cn/models/iic/CosyVoice-300M/summary"), Type: "audio", Parameters: ptr("300M"), Publisher: ptr("Alibaba")},
		{Name: ptr("ChatTTS"), Description: ptr("ChatTTS文本转语音模型"), Url: ptr("https://modelscope.cn/models/pzc163/chatTTS/summary"), Type: "audio", Parameters: ptr("Unknown"), Publisher: ptr("2Noise")},
		{Name: ptr("Stable-Diffusion-3-Medium"), Description: ptr("Stability AI文生图模型"), Url: ptr("https://modelscope.cn/models/AI-ModelScope/stable-diffusion-3-medium-diffusers/summary"), Type: "vision", Parameters: ptr("Medium"), Publisher: ptr("Stability AI")},
		{Name: ptr("Kolors"), Description: ptr("快手可图文生图模型"), Url: ptr("https://modelscope.cn/models/Kwai-Kolors/Kolors/summary"), Type: "vision", Parameters: ptr("Unknown"), Publisher: ptr("Kuaishou")},
		{Name: ptr("Flux.1-dev"), Description: ptr("Black Forest Labs文生图模型"), Url: ptr("https://modelscope.cn/models/AI-ModelScope/FLUX.1-dev/summary"), Type: "vision", Parameters: ptr("12B"), Publisher: ptr("Black Forest Labs")},
		{Name: ptr("CodeQwen1.5-7B-Chat"), Description: ptr("阿里云通义千问代码生成模型"), Url: ptr("https://modelscope.cn/models/qwen/CodeQwen1.5-7B-Chat/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("Qwen")},
		{Name: ptr("StarCoder2-15B"), Description: ptr("BigCode代码生成模型"), Url: ptr("https://modelscope.cn/models/AI-ModelScope/starcoder2-15b/summary"), Type: "general_llm", Parameters: ptr("15B"), Publisher: ptr("BigCode")},
		{Name: ptr("DeepSeek-Prover-V1.5-RL"), Description: ptr("深度求索数学证明模型"), Url: ptr("https://modelscope.cn/models/deepseek-ai/DeepSeek-Prover-V1.5-RL/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("DeepSeek")},
		{Name: ptr("MathGLM"), Description: ptr("智谱AI数学模型"), Url: ptr("https://modelscope.cn/models/ZhipuAI/MathGLM/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ZhipuAI")},
		{Name: ptr("AquilaChat2-34B"), Description: ptr("悟道·天鹰对话模型"), Url: ptr("https://modelscope.cn/models/BAAI/AquilaChat2-34B/summary"), Type: "general_llm", Parameters: ptr("34B"), Publisher: ptr("BAAI")},
		{Name: ptr("Yuan2.0-102B"), Description: ptr("浪潮源2.0大模型"), Url: ptr("https://modelscope.cn/models/IEIT-Yuan/Yuan2-102B-hf/summary"), Type: "general_llm", Parameters: ptr("102B"), Publisher: ptr("IEIT")},
		{Name: ptr("Skywork-13B-Base"), Description: ptr("昆仑万维天工大模型"), Url: ptr("https://modelscope.cn/models/Skywork/Skywork-13B-Base/summary"), Type: "general_llm", Parameters: ptr("13B"), Publisher: ptr("Kunlun")},
		{Name: ptr("Qwen-Max"), Description: ptr("阿里云通义千问Max版本"), Url: ptr("https://modelscope.cn/models/qwen/Qwen-Max/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Qwen")},
		{Name: ptr("Qwen-Plus"), Description: ptr("阿里云通义千问Plus版本"), Url: ptr("https://modelscope.cn/models/qwen/Qwen-Plus/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Qwen")},
		{Name: ptr("Qwen-Turbo"), Description: ptr("阿里云通义千问Turbo版本"), Url: ptr("https://modelscope.cn/models/qwen/Qwen-Turbo/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Qwen")},
		{Name: ptr("Baichuan-53B"), Description: ptr("百川智能530亿参数模型"), Url: ptr("https://modelscope.cn/models/baichuan-inc/Baichuan-53B/summary"), Type: "general_llm", Parameters: ptr("53B"), Publisher: ptr("Baichuan")},
		{Name: ptr("Baichuan-NPC-Turbo"), Description: ptr("百川智能角色扮演模型"), Url: ptr("https://modelscope.cn/models/baichuan-inc/Baichuan-NPC-Turbo/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Baichuan")},
		{Name: ptr("ChatGLM-Turbo"), Description: ptr("智谱AI Turbo版本"), Url: ptr("https://modelscope.cn/models/ZhipuAI/ChatGLM-Turbo/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ZhipuAI")},
		{Name: ptr("ChatGLM-Pro"), Description: ptr("智谱AI Pro版本"), Url: ptr("https://modelscope.cn/models/ZhipuAI/ChatGLM-Pro/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ZhipuAI")},
		{Name: ptr("ChatGLM-Std"), Description: ptr("智谱AI Std版本"), Url: ptr("https://modelscope.cn/models/ZhipuAI/ChatGLM-Std/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ZhipuAI")},
		{Name: ptr("ChatGLM-Lite"), Description: ptr("智谱AI Lite版本"), Url: ptr("https://modelscope.cn/models/ZhipuAI/ChatGLM-Lite/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ZhipuAI")},
		{Name: ptr("MiniMax-abab6"), Description: ptr("MiniMax abab6模型"), Url: ptr("https://modelscope.cn/models/MiniMax/abab6/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("MiniMax")},
		{Name: ptr("MiniMax-abab5.5"), Description: ptr("MiniMax abab5.5模型"), Url: ptr("https://modelscope.cn/models/MiniMax/abab5.5/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("MiniMax")},
		{Name: ptr("Moonshot-v1-8k"), Description: ptr("月之暗面8k上下文模型"), Url: ptr("https://modelscope.cn/models/Moonshot/moonshot-v1-8k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Moonshot AI")},
		{Name: ptr("Moonshot-v1-32k"), Description: ptr("月之暗面32k上下文模型"), Url: ptr("https://modelscope.cn/models/Moonshot/moonshot-v1-32k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Moonshot AI")},
		{Name: ptr("Moonshot-v1-128k"), Description: ptr("月之暗面128k上下文模型"), Url: ptr("https://modelscope.cn/models/Moonshot/moonshot-v1-128k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Moonshot AI")},
		{Name: ptr("Doubao-Pro-4k"), Description: ptr("字节跳动豆包Pro 4k"), Url: ptr("https://modelscope.cn/models/ByteDance/Doubao-Pro-4k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ByteDance")},
		{Name: ptr("Doubao-Pro-32k"), Description: ptr("字节跳动豆包Pro 32k"), Url: ptr("https://modelscope.cn/models/ByteDance/Doubao-Pro-32k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ByteDance")},
		{Name: ptr("Doubao-Pro-128k"), Description: ptr("字节跳动豆包Pro 128k"), Url: ptr("https://modelscope.cn/models/ByteDance/Doubao-Pro-128k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ByteDance")},
		{Name: ptr("Doubao-Lite-4k"), Description: ptr("字节跳动豆包Lite 4k"), Url: ptr("https://modelscope.cn/models/ByteDance/Doubao-Lite-4k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ByteDance")},
		{Name: ptr("Doubao-Lite-32k"), Description: ptr("字节跳动豆包Lite 32k"), Url: ptr("https://modelscope.cn/models/ByteDance/Doubao-Lite-32k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ByteDance")},
		{Name: ptr("Doubao-Lite-128k"), Description: ptr("字节跳动豆包Lite 128k"), Url: ptr("https://modelscope.cn/models/ByteDance/Doubao-Lite-128k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("ByteDance")},
		{Name: ptr("Hunyuan-Pro"), Description: ptr("腾讯混元Pro"), Url: ptr("https://modelscope.cn/models/Tencent/Hunyuan-Pro/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Tencent")},
		{Name: ptr("Hunyuan-Standard"), Description: ptr("腾讯混元Standard"), Url: ptr("https://modelscope.cn/models/Tencent/Hunyuan-Standard/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Tencent")},
		{Name: ptr("Hunyuan-Lite"), Description: ptr("腾讯混元Lite"), Url: ptr("https://modelscope.cn/models/Tencent/Hunyuan-Lite/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Tencent")},
		{Name: ptr("Spark-V4.0"), Description: ptr("讯飞星火4.0"), Url: ptr("https://modelscope.cn/models/iFLYTEK/Spark-V4.0/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("iFLYTEK")},
		{Name: ptr("Spark-V3.5"), Description: ptr("讯飞星火3.5"), Url: ptr("https://modelscope.cn/models/iFLYTEK/Spark-V3.5/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("iFLYTEK")},
		{Name: ptr("Spark-V3.0"), Description: ptr("讯飞星火3.0"), Url: ptr("https://modelscope.cn/models/iFLYTEK/Spark-V3.0/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("iFLYTEK")},
		{Name: ptr("SenseChat-5.0"), Description: ptr("商汤日日新5.0"), Url: ptr("https://modelscope.cn/models/SenseTime/SenseChat-5.0/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("SenseTime")},
		{Name: ptr("SenseChat-V4"), Description: ptr("商汤日日新4.0"), Url: ptr("https://modelscope.cn/models/SenseTime/SenseChat-V4/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("SenseTime")},
		{Name: ptr("Step-2-16k"), Description: ptr("阶跃星辰Step-2 16k"), Url: ptr("https://modelscope.cn/models/StepFun/Step-2-16k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("StepFun")},
		{Name: ptr("Step-1-32k"), Description: ptr("阶跃星辰Step-1 32k"), Url: ptr("https://modelscope.cn/models/StepFun/Step-1-32k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("StepFun")},
		{Name: ptr("Step-1-8k"), Description: ptr("阶跃星辰Step-1 8k"), Url: ptr("https://modelscope.cn/models/StepFun/Step-1-8k/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("StepFun")},
		{Name: ptr("Yi-Large"), Description: ptr("零一万物Yi-Large"), Url: ptr("https://modelscope.cn/models/01ai/Yi-Large/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("01.AI")},
		{Name: ptr("Yi-Medium"), Description: ptr("零一万物Yi-Medium"), Url: ptr("https://modelscope.cn/models/01ai/Yi-Medium/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("01.AI")},
		{Name: ptr("Yi-Spark"), Description: ptr("零一万物Yi-Spark"), Url: ptr("https://modelscope.cn/models/01ai/Yi-Spark/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("01.AI")},
		{Name: ptr("DeepSeek-V2"), Description: ptr("深度求索V2"), Url: ptr("https://modelscope.cn/models/deepseek-ai/DeepSeek-V2/summary"), Type: "general_llm", Parameters: ptr("236B"), Publisher: ptr("DeepSeek")},
		{Name: ptr("DeepSeek-67B-Chat"), Description: ptr("深度求索67亿参数对话模型"), Url: ptr("https://modelscope.cn/models/deepseek-ai/deepseek-llm-67b-chat/summary"), Type: "general_llm", Parameters: ptr("67B"), Publisher: ptr("DeepSeek")},
		{Name: ptr("DeepSeek-7B-Chat"), Description: ptr("深度求索7亿参数对话模型"), Url: ptr("https://modelscope.cn/models/deepseek-ai/deepseek-llm-7b-chat/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("DeepSeek")},
		{Name: ptr("InternLM2-20B-Chat"), Description: ptr("书生·浦语2.0版本200亿参数对话模型"), Url: ptr("https://modelscope.cn/models/Shanghai_AI_Laboratory/internlm2-chat-20b/summary"), Type: "general_llm", Parameters: ptr("20B"), Publisher: ptr("Shanghai AI Lab")},
		{Name: ptr("InternLM2-7B-Chat"), Description: ptr("书生·浦语2.0版本70亿参数对话模型"), Url: ptr("https://modelscope.cn/models/Shanghai_AI_Laboratory/internlm2-chat-7b/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("Shanghai AI Lab")},
		{Name: ptr("Baichuan2-Turbo"), Description: ptr("百川智能Turbo"), Url: ptr("https://modelscope.cn/models/baichuan-inc/Baichuan2-Turbo/summary"), Type: "general_llm", Parameters: ptr("Unknown"), Publisher: ptr("Baichuan")},
		{Name: ptr("Baichuan2-53B"), Description: ptr("百川智能530亿参数"), Url: ptr("https://modelscope.cn/models/baichuan-inc/Baichuan2-53B/summary"), Type: "general_llm", Parameters: ptr("53B"), Publisher: ptr("Baichuan")},
		{Name: ptr("TeleChat-52B"), Description: ptr("中国电信星辰大模型52B"), Url: ptr("https://modelscope.cn/models/TeleAI/TeleChat-52B/summary"), Type: "general_llm", Parameters: ptr("52B"), Publisher: ptr("TeleAI")},
		{Name: ptr("TeleChat-12B-v2"), Description: ptr("中国电信星辰大模型12B v2"), Url: ptr("https://modelscope.cn/models/TeleAI/TeleChat-12B-v2/summary"), Type: "general_llm", Parameters: ptr("12B"), Publisher: ptr("TeleAI")},
		{Name: ptr("TeleChat-7B"), Description: ptr("中国电信星辰大模型7B"), Url: ptr("https://modelscope.cn/models/TeleAI/TeleChat-7B/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("TeleAI")},
		{Name: ptr("TigerBot-70B-Chat"), Description: ptr("虎博科技70B对话模型"), Url: ptr("https://modelscope.cn/models/TigerResearch/tigerbot-70b-chat/summary"), Type: "general_llm", Parameters: ptr("70B"), Publisher: ptr("TigerResearch")},
		{Name: ptr("TigerBot-13B-Chat"), Description: ptr("虎博科技13B对话模型"), Url: ptr("https://modelscope.cn/models/TigerResearch/tigerbot-13b-chat/summary"), Type: "general_llm", Parameters: ptr("13B"), Publisher: ptr("TigerResearch")},
		{Name: ptr("TigerBot-7B-Chat"), Description: ptr("虎博科技7B对话模型"), Url: ptr("https://modelscope.cn/models/TigerResearch/tigerbot-7b-chat/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("TigerResearch")},
		{Name: ptr("MOSS-Moon-003-sft"), Description: ptr("复旦MOSS模型"), Url: ptr("https://modelscope.cn/models/fnlp/moss-moon-003-sft/summary"), Type: "general_llm", Parameters: ptr("16B"), Publisher: ptr("FudanNLP")},
		{Name: ptr("CPM-Bee-10B"), Description: ptr("OpenBMB CPM-Bee 10B"), Url: ptr("https://modelscope.cn/models/OpenBMB/cpm-bee-10b/summary"), Type: "general_llm", Parameters: ptr("10B"), Publisher: ptr("OpenBMB")},
		{Name: ptr("CPM-Bee-5B"), Description: ptr("OpenBMB CPM-Bee 5B"), Url: ptr("https://modelscope.cn/models/OpenBMB/cpm-bee-5b/summary"), Type: "general_llm", Parameters: ptr("5B"), Publisher: ptr("OpenBMB")},
		{Name: ptr("CPM-Bee-2B"), Description: ptr("OpenBMB CPM-Bee 2B"), Url: ptr("https://modelscope.cn/models/OpenBMB/cpm-bee-2b/summary"), Type: "general_llm", Parameters: ptr("2B"), Publisher: ptr("OpenBMB")},
		{Name: ptr("Aquila-33B-Chat"), Description: ptr("悟道·天鹰33B对话模型"), Url: ptr("https://modelscope.cn/models/BAAI/Aquila-33B-Chat/summary"), Type: "general_llm", Parameters: ptr("33B"), Publisher: ptr("BAAI")},
		{Name: ptr("Aquila-7B-Chat"), Description: ptr("悟道·天鹰7B对话模型"), Url: ptr("https://modelscope.cn/models/BAAI/Aquila-7B-Chat/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("BAAI")},
		{Name: ptr("Linly-Chinese-LLaMA-2-7B"), Description: ptr("Linly中文LLaMA2 7B"), Url: ptr("https://modelscope.cn/models/CVI-SZU/Linly-Chinese-LLaMA-2-7b/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("CVI-SZU")},
		{Name: ptr("Linly-Chinese-LLaMA-2-13B"), Description: ptr("Linly中文LLaMA2 13B"), Url: ptr("https://modelscope.cn/models/CVI-SZU/Linly-Chinese-LLaMA-2-13b/summary"), Type: "general_llm", Parameters: ptr("13B"), Publisher: ptr("CVI-SZU")},
		{Name: ptr("Chinese-Alpaca-2-7B"), Description: ptr("中文Alpaca 2 7B"), Url: ptr("https://modelscope.cn/models/ziqingyang/chinese-alpaca-2-7b/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("Yiming Cui")},
		{Name: ptr("Chinese-Alpaca-2-13B"), Description: ptr("中文Alpaca 2 13B"), Url: ptr("https://modelscope.cn/models/ziqingyang/chinese-alpaca-2-13b/summary"), Type: "general_llm", Parameters: ptr("13B"), Publisher: ptr("Yiming Cui")},
		{Name: ptr("Chinese-Llama-2-7B"), Description: ptr("中文Llama 2 7B"), Url: ptr("https://modelscope.cn/models/ziqingyang/chinese-llama-2-7b/summary"), Type: "general_llm", Parameters: ptr("7B"), Publisher: ptr("Yiming Cui")},
		{Name: ptr("Chinese-Llama-2-13B"), Description: ptr("中文Llama 2 13B"), Url: ptr("https://modelscope.cn/models/ziqingyang/chinese-llama-2-13b/summary"), Type: "general_llm", Parameters: ptr("13B"), Publisher: ptr("Yiming Cui")},
	}
	if err := global.GVA_DB.Create(&data).Error; err != nil {
		zap.L().Error("初始化大模型数据失败", zap.Error(err))
	}
}
