<template>
  <div>
    <div class="echart"
         id="echart-line"
         :style="{float:'left',width: '100%', height: '460px'}"></div>
    <input type="text"
           placeholder="开始时间: 2h"
           v-model="formMess.time_begin">
    <span> ~ </span>
    <input type="text"
           placeholder="结束时间: 0h"
           v-model="formMess.time_end">
    <span class="button"
          @click="onSubmit()">提交</span>
  </div>
</template>

<script>

import * as echarts from 'echarts'
import axios from "axios"

export default {
  name: 'ChartPprof',
  data() {
    return {
      formMess: {
        "time_begin": '',
        "time_end": '',
      },
    }
  },
  methods: {
    initChart(xAxis, yAxis, label) {
      let chartDom = document.getElementById('echart-line')
      let myChart = echarts.init(chartDom)
      let option = {
        tooltip: {
          trigger: 'axis',
          position: function (pt) {
            return [pt[0], '10%']
          }
        },
        title: {
          left: 'center',
          text: '进程内存占用',
        },
        toolbox: {
          feature: {
            magicType: {
              type: ['line', 'bar']
            },
            dataZoom: {
              yAxisIndex: 'none'
            },
            restore: {},
            saveAsImage: {}
          }
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: xAxis
        },
        yAxis: {
          name: label,
          type: 'value',
          boundaryGap: [0, '100%']
        },
        dataZoom: [{
          type: 'inside',
          start: 0,
          end: 10
        }, {
          start: 0,
          end: 10
        }],
        series: [
          {
            name: 'proc_mem_rss',
            type: 'line',
            symbol: 'none',
            sampling: 'lttb',
            itemStyle: {
              color: 'rgb(255, 70, 131)'
            },
            areaStyle: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [{
                offset: 0,
                color: 'rgb(255, 158, 68)'
              }, {
                offset: 1,
                color: 'rgb(255, 70, 131)'
              }])
            },
            data: yAxis
          }
        ]
      }

      option && myChart.setOption(option)
      myChart.on('click', function (param) {
        window.open("http://localhost:10000/ui/peek?file=" + param.data.url)
      })
    },
    formatDate(value) {
      let date = new Date(value * 1000)
      let y = date.getFullYear()
      let MM = date.getMonth() + 1
      MM = MM < 10 ? "0" + MM : MM
      let d = date.getDate()
      d = d < 10 ? "0" + d : d
      let h = date.getHours()
      h = h < 10 ? "0" + h : h
      let m = date.getMinutes()
      m = m < 10 ? "0" + m : m
      let s = date.getSeconds()
      s = s < 10 ? "0" + s : s
      return y + "-" + MM + "-" + d + " " + h + ":" + m + ":" + s
    },
    setChart(time_begin, time_end) {
      axios
        .get('http://127.0.0.1:10000/ui/log?start=' + time_begin + '&' + 'end=' + time_end)
        .then(response => {
          console.log(response.data)
          let xAxis = []
          let yAxis = []
          let unit = ''
          if (response.data.length === 0) {
            alert('No data retrieved!')
            return
          }
          response.data.forEach(mem => {
            if (unit === '' && mem["unit"] !== 'B') {
              unit = mem["unit"]
            }
            let a = { url: '', value: 0 }
            xAxis[xAxis.length] = this.formatDate(mem["created_at"])
            if (mem["file_path"] !== undefined) {
              a.url = mem["file_path"]
            }
            if (mem["unit"] === 'B') {
              a.value = 0
            } else if (unit === 'kB') {
              if (mem["unit"] === 'kB') {
                a.value = mem["pro_mem_rss"]
              } else if (mem["unit"] === 'MB') {
                a.value = mem["pro_mem_rss"] * 1024
              } else if (mem["unit"] === 'GB') {
                a.value = mem["pro_mem_rss"] * 1024 * 1024
              }
            } else if (unit === 'MB') {
              if (mem["unit"] === 'kB') {
                a.value = mem["pro_mem_rss"] / 1024
              } else if (mem["unit"] === 'MB') {
                a.value = mem["pro_mem_rss"]
              } else if (mem["unit"] === 'GB') {
                a.value = mem["pro_mem_rss"] * 1024
              }
            } else if (unit === 'GB') {
              if (mem["unit"] === 'kB') {
                a.value = mem["pro_mem_rss"] / 1024 / 1024
              } else if (mem["unit"] === 'MB') {
                a.value = mem["pro_mem_rss"] / 1024
              } else if (mem["unit"] === 'GB') {
                a.value = mem["pro_mem_rss"]
              }
            }
            yAxis[yAxis.length] = a
          })
          this.initChart(xAxis, yAxis, unit)
        })
    },
    onSubmit() {
      this.setChart(this.formMess["time_begin"], this.formMess["time_end"])
    },
  },
  mounted() {
    this.setChart('2h')
  }

}
</script>

<style>
.button {
  background-color: #4caf50; /* Green */
  border: none;
  color: white;
  padding: 6px 8px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
}

.button:hover {
  color: white;
  padding: 6px 8px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  background-color: #ff69b4;
  border-color: #ff69b4;
}

.button:focus {
  color: white;
  padding: 6px 8px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
  background-color: #ff69b4;
  border-color: #ff69b4;
}
</style>
