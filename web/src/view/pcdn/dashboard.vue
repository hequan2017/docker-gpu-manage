<template>
  <div class="pcdn-page">
    <el-row :gutter="16" class="mb-4">
      <el-col :xs="24" :md="8">
        <el-card shadow="hover">
          <div class="metric-title">实时带宽（Gbps）</div>
          <div class="metric-value">{{ summary.realtimeBandwidth }}</div>
        </el-card>
      </el-col>
      <el-col :xs="24" :md="8">
        <el-card shadow="hover">
          <div class="metric-title">命中率</div>
          <div class="metric-value">{{ summary.hitRate }}%</div>
        </el-card>
      </el-col>
      <el-col :xs="24" :md="8">
        <el-card shadow="hover">
          <div class="metric-title">调度成功率</div>
          <div class="metric-value">{{ summary.dispatchSuccessRate }}%</div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="16">
      <el-col :xs="24" :md="12">
        <el-card shadow="never">
          <template #header>实时带宽趋势</template>
          <charts :options="bandwidthChartOption" height="280px" />
        </el-card>
      </el-col>
      <el-col :xs="24" :md="12">
        <el-card shadow="never">
          <template #header>命中率与调度成功率</template>
          <charts :options="rateChartOption" height="280px" />
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import charts from '@/components/charts/index.vue'
import useChartOption from '@/hooks/charts'
import { getPcdnDashboardMetrics } from '@/api/pcdn'

defineOptions({ name: 'PcdnDashboard' })

const summary = ref({
  realtimeBandwidth: 18.7,
  hitRate: 92.8,
  dispatchSuccessRate: 97.6
})

const xAxisData = ref(['10:00', '10:05', '10:10', '10:15', '10:20', '10:25', '10:30'])
const bandwidthData = ref([12.1, 13.4, 14.8, 15.6, 17.2, 18.4, 18.7])
const hitRateData = ref([89.2, 90.1, 91.5, 92.4, 92.8, 92.6, 92.8])
const dispatchRateData = ref([95.4, 95.9, 96.3, 96.8, 97.1, 97.5, 97.6])

const { chartOption: bandwidthChartOption } = useChartOption((isDark) => {
  const color = isDark ? '#ddd' : '#303133'
  return {
    tooltip: { trigger: 'axis' },
    xAxis: { type: 'category', data: xAxisData.value, axisLabel: { color } },
    yAxis: { type: 'value', axisLabel: { color } },
    series: [{ data: bandwidthData.value, type: 'line', smooth: true, areaStyle: {} }]
  }
})

const { chartOption: rateChartOption } = useChartOption((isDark) => {
  const color = isDark ? '#ddd' : '#303133'
  return {
    tooltip: { trigger: 'axis' },
    legend: { data: ['命中率', '调度成功率'], textStyle: { color } },
    xAxis: { type: 'category', data: xAxisData.value, axisLabel: { color } },
    yAxis: { type: 'value', axisLabel: { color } },
    series: [
      { name: '命中率', data: hitRateData.value, type: 'line', smooth: true },
      { name: '调度成功率', data: dispatchRateData.value, type: 'line', smooth: true }
    ]
  }
})

const updateMetrics = (data) => {
  summary.value = {
    realtimeBandwidth: data?.realtimeBandwidth ?? summary.value.realtimeBandwidth,
    hitRate: data?.hitRate ?? summary.value.hitRate,
    dispatchSuccessRate: data?.dispatchSuccessRate ?? summary.value.dispatchSuccessRate
  }
  xAxisData.value = data?.timeline || xAxisData.value
  bandwidthData.value = data?.bandwidthTrend || bandwidthData.value
  hitRateData.value = data?.hitRateTrend || hitRateData.value
  dispatchRateData.value = data?.dispatchSuccessTrend || dispatchRateData.value
}

const loadMetrics = async () => {
  try {
    const res = await getPcdnDashboardMetrics()
    updateMetrics(res?.data)
  } catch (e) {
    updateMetrics({})
  }
}

onMounted(loadMetrics)
</script>

<style scoped lang="scss">
.pcdn-page {
  padding: 20px;
}

.mb-4 {
  margin-bottom: 16px;
}

.metric-title {
  font-size: 14px;
  color: #606266;
}

.metric-value {
  margin-top: 10px;
  font-size: 32px;
  font-weight: 700;
  color: #303133;
}
</style>
