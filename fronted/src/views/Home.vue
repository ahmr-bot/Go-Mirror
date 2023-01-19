<!-- eslint-disable vue/no-v-text-v-html-on-component -->
<template>
  <v-app id="inspire">
    <v-main class="bg-grey-lighten-3">
      <v-container>
        <v-row>
          <v-col cols="5">
            <v-card>
      <v-card-title class="text-h6 text-md-h5 text-lg-h4">Welcome!</v-card-title>
      <v-card-text>
        欢迎来到 镜缘 Minecraft 镜像站！<br />
        本站提供 Minecraft 服务端的镜像下载服务。<br />
              后端及前端程序 由 Aehxy 开发，服务端同步工具 由 kingc 开发。
              <br /> 
              网络部分由 PH 提供支持。<br />
              切换服务器:
      </v-card-text>
      <a href="/?server=defaultServer">默认服务器(无CDN)</a><br />
      <a href="/?server=AutoSel">自动选择(CDN)</a><br />
      <a href="/?server=bjServer">北京服务器</a><br />
      <a href="/?server=gzServer">贵州服务器</a><br />
      <a href="/?server=gsServer">甘肃服务器</a><br />
    </v-card>
          </v-col>

          <v-col>
            <v-sheet
              min-height="70vh"
              rounded="lg"
            >
            <v-list density="compact">
      <v-breadcrumbs :items="items">
    <template v-slot:title="{ item }">
      {{ item.title.toUpperCase() }}
    </template>
  </v-breadcrumbs>

      <v-list-item
        v-for="(list,i) in lists"
        :key="i"
        :value="list"
        active-color="primary"
        @click="router.push({path: list.link})"
      >
      <v-avatar>
      <v-img :src="list.img"></v-img>
    </v-avatar>
    {{ list.name  }}
      </v-list-item>
    </v-list>
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
axios({
  method:'get',
  url:server.value + 'list/'
}).then((res)=>{
  console.log(res.data)
  lists.value = res.data.directories
   downlists.value = res.data.directories.map((item)=>{
     return item
   })
   imglists.value = res.data.directories.map((item)=>{
     return "https://cdn.5-5.site/mirror/imgs/" + item + ".png"
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
