<template>
  <div class="ai-chat-container">
    <!-- 侧边栏：会话列表 -->
    <div class="chat-sidebar">
      <div class="sidebar-header">
        <h3>AI 对话</h3>
        <el-button type="primary" icon="Plus" @click="createNewChat">新建对话</el-button>
      </div>
      <div class="conversation-list">
        <div
          v-for="conv in conversations"
          :key="conv.ID"
          :class="['conversation-item', { active: currentConversation?.ID === conv.ID }]"
          @click="selectConversation(conv)"
        >
          <div class="conv-title">{{ conv.title || '新对话' }}</div>
          <div class="conv-info">
            <span class="conv-model">{{ conv.model }}</span>
            <span class="conv-time">{{ formatTime(conv.CreatedAt) }}</span>
          </div>
          <el-button
            type="danger"
            icon="Delete"
            size="small"
            link
            @click.stop="deleteConversation(conv)"
          >
            删除
          </el-button>
        </div>
      </div>
    </div>

    <!-- 主聊天区域 -->
    <div class="chat-main">
      <!-- 消息列表 -->
      <div class="message-container" ref="messageContainer">
        <div v-if="messages.length === 0" class="empty-state">
          <el-empty description="开始新的对话吧！"></el-empty>
        </div>
        <div
          v-for="msg in messages"
          :key="msg.ID"
          :class="['message-item', `message-${msg.role}`]"
        >
          <div class="message-avatar">
            <el-avatar v-if="msg.role === 'user'" :icon="UserFilled" />
            <el-avatar v-else :icon="ChatDotSquare" />
          </div>
          <div class="message-content">
            <div class="message-text">{{ msg.content }}</div>
            <div v-if="msg.tokenCount" class="message-meta">
              Tokens: {{ msg.tokenCount }}
            </div>
          </div>
        </div>
        <!-- 加载状态 -->
        <div v-if="loading" class="message-item message-assistant">
          <div class="message-avatar">
            <el-avatar :icon="ChatDotSquare" />
          </div>
          <div class="message-content">
            <div class="loading-text">
              <span>AI正在思考</span>
              <el-icon class="is-loading"><Loading /></el-icon>
            </div>
          </div>
        </div>
      </div>

      <!-- 输入区域 -->
      <div class="input-container">
        <el-input
          v-model="inputMessage"
          type="textarea"
          :rows="3"
          placeholder="输入消息..."
          @keydown.enter.ctrl="sendMessage"
          :disabled="loading"
        />
        <div class="input-actions">
          <span class="input-hint">Ctrl + Enter 发送</span>
          <el-button
            type="primary"
            icon="Promotion"
            @click="sendMessage"
            :loading="loading"
            :disabled="!inputMessage.trim()"
          >
            发送
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, nextTick, computed } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import { UserFilled, ChatDotSquare, Loading } from '@element-plus/icons-vue'
import {
  sendMessage as sendMessageApi,
  getConversationList,
  createConversation,
  deleteConversation as deleteConversationApi,
  setConversationActive,
  getActiveConversation,
  getMessageList
} from '@/plugin/aiagent/api/conversation'
import { getMessageList as getMessagesApi } from '@/plugin/aiagent/api/message'
import { formatDate } from '@/utils/format'

defineOptions({
  name: 'AIAgentChat'
})

const conversations = ref([])
const currentConversation = ref(null)
const messages = ref([])
const inputMessage = ref('')
const loading = ref(false)
const messageContainer = ref(null)

// 加载会话列表
const loadConversations = async () => {
  const res = await getConversationList({ page: 1, pageSize: 100 })
  if (res.code === 0) {
    conversations.value = res.data.list || []
  }
}

// 加载当前会话的消息
const loadMessages = async () => {
  if (!currentConversation.value) return
  const res = await getMessagesApi({
    conversationID: currentConversation.value.ID,
    page: 1,
    pageSize: 1000
  })
  if (res.code === 0) {
    messages.value = res.data.list || []
    scrollToBottom()
  }
}

// 选择会话
const selectConversation = async (conv) => {
  currentConversation.value = conv
  await loadMessages()
}

// 新建对话
const createNewChat = async () => {
  // 创建新会话但不发送消息
  const res = await createConversation({
    title: '新对话',
    model: 'glm-4-plus',
    temperature: 0.7,
    maxTokens: 4096,
    isActive: true
  })
  if (res.code === 0) {
    await loadConversations()
    await selectConversation(res.data)
  }
}

// 删除会话
const deleteConversation = async (conv) => {
  ElMessageBox.confirm('确定要删除这个对话吗？', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async () => {
    const res = await deleteConversationApi({ ID: conv.ID })
    if (res.code === 0) {
      ElMessage.success('删除成功')
      if (currentConversation.value?.ID === conv.ID) {
        currentConversation.value = null
        messages.value = []
      }
      await loadConversations()
    }
  })
}

// 发送消息
const sendMessage = async () => {
  const message = inputMessage.value.trim()
  if (!message || loading.value) return

  loading.value = true
  inputMessage.value = ''

  try {
    const res = await sendMessageApi({
      conversationID: currentConversation.value?.ID,
      message: message
    })

    if (res.code === 0) {
      // 如果是新的会话，更新当前会话
      if (res.data.conversationID && (!currentConversation.value || currentConversation.value.ID !== res.data.conversationID)) {
        currentConversation.value = { ID: res.data.conversationID }
        await loadConversations()
      }

      // 添加用户消息和AI回复到列表
      messages.value.push({
        ID: Date.now(),
        role: 'user',
        content: message
      })
      messages.value.push({
        ID: res.data.messageID,
        role: 'assistant',
        content: res.data.content,
        tokenCount: res.data.usage.totalTokens
      })

      scrollToBottom()
    }
  } catch (error) {
    ElMessage.error('发送消息失败: ' + error.message)
  } finally {
    loading.value = false
  }
}

// 滚动到底部
const scrollToBottom = () => {
  nextTick(() => {
    if (messageContainer.value) {
      messageContainer.value.scrollTop = messageContainer.value.scrollHeight
    }
  })
}

// 格式化时间
const formatTime = (time) => {
  if (!time) return ''
  const date = new Date(time)
  const now = new Date()
  const diff = now - date

  if (diff < 60000) return '刚刚'
  if (diff < 3600000) return `${Math.floor(diff / 60000)}分钟前`
  if (diff < 86400000) return `${Math.floor(diff / 3600000)}小时前`
  return formatDate(time)
}

// 初始化
onMounted(async () => {
  await loadConversations()

  // 尝试获取激活的会话
  const activeRes = await getActiveConversation()
  if (activeRes.code === 0 && activeRes.data) {
    currentConversation.value = activeRes.data
    await loadMessages()
  }
})
</script>

<style scoped>
.ai-chat-container {
  display: flex;
  height: calc(100vh - 100px);
  background: #f5f7fa;
}

.chat-sidebar {
  width: 300px;
  background: #fff;
  border-right: 1px solid #e4e7ed;
  display: flex;
  flex-direction: column;
}

.sidebar-header {
  padding: 20px;
  border-bottom: 1px solid #e4e7ed;
}

.sidebar-header h3 {
  margin: 0 0 10px 0;
  font-size: 18px;
}

.conversation-list {
  flex: 1;
  overflow-y: auto;
}

.conversation-item {
  padding: 15px 20px;
  cursor: pointer;
  border-bottom: 1px solid #f5f7fa;
  transition: background 0.2s;
}

.conversation-item:hover {
  background: #f5f7fa;
}

.conversation-item.active {
  background: #e6f7ff;
  border-left: 3px solid #1890ff;
}

.conv-title {
  font-weight: 500;
  margin-bottom: 5px;
}

.conv-info {
  display: flex;
  justify-content: space-between;
  font-size: 12px;
  color: #909399;
}

.chat-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: #fff;
}

.message-container {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
}

.empty-state {
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.message-item {
  display: flex;
  margin-bottom: 20px;
}

.message-user {
  flex-direction: row-reverse;
}

.message-avatar {
  margin: 0 10px;
}

.message-content {
  max-width: 70%;
}

.message-user .message-content {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
}

.message-text {
  padding: 10px 15px;
  border-radius: 8px;
  line-height: 1.5;
  word-break: break-word;
}

.message-user .message-text {
  background: #1890ff;
  color: #fff;
}

.message-assistant .message-text {
  background: #f5f7fa;
  color: #333;
}

.message-meta {
  font-size: 12px;
  color: #909399;
  margin-top: 5px;
}

.loading-text {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 15px;
  background: #f5f7fa;
  border-radius: 8px;
}

.input-container {
  padding: 20px;
  border-top: 1px solid #e4e7ed;
}

.input-actions {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 10px;
}

.input-hint {
  font-size: 12px;
  color: #909399;
}
</style>
