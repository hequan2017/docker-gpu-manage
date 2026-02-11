<template>
  <el-dialog
    v-model="visible"
    :title="`终端 - ${podName} (${containerName})`"
    width="80%"
    :before-close="handleClose"
    top="5vh"
    destroy-on-close
  >
    <div ref="terminalContainer" class="terminal-container"></div>
  </el-dialog>
</template>

<script setup>
import { ref, onMounted, onBeforeUnmount, watch, nextTick } from 'vue'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import 'xterm/css/xterm.css'
import { useUserStore } from '@/pinia/modules/user'

const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  clusterName: {
    type: String,
    required: true
  },
  namespace: {
    type: String,
    required: true
  },
  podName: {
    type: String,
    required: true
  },
  containerName: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue', 'close'])

const visible = ref(props.modelValue)
const terminalContainer = ref(null)
const userStore = useUserStore()

let term = null
let socket = null
let fitAddon = null

watch(() => props.modelValue, (val) => {
  visible.value = val
  if (val) {
    nextTick(() => {
      initTerminal()
    })
  } else {
    closeTerminal()
  }
})

watch(() => visible.value, (val) => {
  emit('update:modelValue', val)
})

const initTerminal = () => {
  if (term) {
    return
  }

  term = new Terminal({
    cursorBlink: true,
    fontSize: 14,
    fontFamily: 'Menlo, Monaco, "Courier New", monospace',
    theme: {
      background: '#1e1e1e',
      foreground: '#ffffff',
    }
  })

  fitAddon = new FitAddon()
  term.loadAddon(fitAddon)
  term.open(terminalContainer.value)
  fitAddon.fit()

  // 连接WebSocket
  // const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws'
  const protocol = 'ws'
  const host = '127.0.0.1:8890'
  // 注意：需要确保后端API路径正确，且包含鉴权token（如果需要）
  const token = userStore.token
  const url = `${protocol}://${host}/k8s/pod/exec?clusterName=${props.clusterName}&namespace=${props.namespace}&podName=${props.podName}&container=${props.containerName}&command=/bin/sh&x-token=${token}`
  
  socket = new WebSocket(url)

  socket.onopen = () => {
    term.write('\r\n\x1b[32mConnected to terminal\x1b[0m\r\n')
    fitAddon.fit()
  }

  socket.onmessage = (event) => {
    term.write(event.data)
  }

  socket.onclose = () => {
    term.write('\r\n\x1b[31mConnection closed\x1b[0m\r\n')
  }

  socket.onerror = (error) => {
    term.write(`\r\n\x1b[31mConnection error: ${error}\x1b[0m\r\n`)
  }

  term.onData((data) => {
    if (socket && socket.readyState === WebSocket.OPEN) {
      socket.send(data)
    }
  })

  // 监听窗口大小变化
  window.addEventListener('resize', handleResize)
}

const handleResize = () => {
  if (fitAddon) {
    fitAddon.fit()
  }
}

const closeTerminal = () => {
  if (socket) {
    socket.close()
    socket = null
  }
  if (term) {
    term.dispose()
    term = null
  }
  window.removeEventListener('resize', handleResize)
}

const handleClose = (done) => {
  closeTerminal()
  emit('close')
  done()
}

onBeforeUnmount(() => {
  closeTerminal()
})
</script>

<style scoped>
.terminal-container {
  height: 60vh;
  width: 100%;
  background-color: #1e1e1e;
  padding: 10px;
  box-sizing: border-box;
  overflow: hidden;
}
</style>
