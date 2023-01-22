<!-- eslint-disable vue/no-v-text-v-html-on-component -->
<template>
  <v-app id="inspire">
    <v-main class="bg-grey-lighten-3">
      <v-container>
        <v-row>
          <v-col>
            <v-card>
      <v-card-title class="text-h6 text-md-h5 text-lg-h4">Welcome!</v-card-title>
      <v-card-text>
        欢迎来到 镜缘 Minecraft 镜像站！<br />
        本站提供 Minecraft 服务端的镜像下载服务。<br />
              后端及前端程序 由 Aehxy 开发，服务端同步工具 由 kingc 开发。
              <br /> 
              网络部分由 PH 提供支持。<br />
              当前选择的存储服务器: {{ servername }}<br />
              手动切换服务器:
      </v-card-text>
      <div class="d-flex justify-center align-baseline">
      <v-btn href="/?server=defaultServer">自动选择服务器</v-btn><br />
      <v-btn href="/?server=bjServer">北京服务器(IPV4)</v-btn><br />
      </div>
      <div class="d-flex justify-center align-baseline">
      <v-btn href="/?server=gzServer">贵州服务器(IPV4)</v-btn><br />
      <v-btn href="/?server=gsServer">甘肃服务器(IPV6)</v-btn><br />
      </div>  
    </v-card>
          </v-col>

          <v-col>
            <v-sheet
              min-height="70vh"
              rounded="lg"
            >
      <v-breadcrumbs :items="items">
    <template v-slot:title="{ item }">
      {{ item.title.toUpperCase() }}
    </template>
  </v-breadcrumbs>
  <v-card
    class="mx-auto"
    max-width="500"
  >
  <v-container fluid>
      <v-row dense>
        <v-col
          v-for="list in lists"
          :key="list.name"
        >
  <v-card min-width="150"  max-width='300' outlined>
            <v-img
              :src="list.img"
              class="align-end"
              :aspect-ratio="1"
              gradient="to bottom, rgba(0,0,0,.1), rgba(0,0,0,.5)"
            >
              <v-card-title class="text-white" v-text="list.name"></v-card-title>
            </v-img>

            <v-card-actions>
              <v-spacer></v-spacer>

              <v-btn size="small" color="surface-variant" variant="text" icon="mdi-heart" @click="router.push({path: list.link})">查看</v-btn>
            </v-card-actions>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-card>
            </v-sheet>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <Copyright />
  </v-app>
</template>

<script setup>
import Copyright from '@/components/Copyright.vue';
import config from '@/config/config';
import router from '@/router';
import axios from 'axios';
import { ref } from 'vue';
const query = router.currentRoute.value.query
// 如果有 server 参数 则使用 server 参数在config.prod 中查找对应的服务器地址 否则先检查是否有本地存储的服务器地址 如果没有则使用默认服务器地址 如果有则使用本地存储的服务器地址
const server = ref(query.server ? config[query.server] : localStorage.getItem('server') ? localStorage.getItem('server') : config.defaultServer)
localStorage.setItem('server',server.value)

// axios 访问API 获取json数据
const items = [
        {
          title: '下载列表',
          disabled: false,
          href: '/',
        },
      ]
const lists = ref([])
const downlists = ref([])
const imglists = ref([])
const servername = ref([])
axios({
  method:'get',
  url:server.value + 'location'
}).then((res)=>{
  servername.value = res.data.server_location
})
axios({
  method:'get',
  url:server.value + 'list/'
}).then((res)=>{
  lists.value = res.data.directories
   downlists.value = res.data.directories.map((item)=>{
     return item
   })
   imglists.value = res.data.directories.map((item)=>{
     return "https://pic.5-5.site/i/2023/01/20/" + item + ".webp"
   })
   // 使list 中同时包含 imglists 和 downlists
     lists.value = lists.value.map((item,index)=>{
       return {
         name:item,
         img:imglists.value[index],
         link:downlists.value[index]
       }
})
})
</script>
