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
                网络部分由 PH 提供支持。
        </v-card-text>
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
  <v-table>
    <thead>
      <tr>
        <th class="text-left">
         文件名
        </th>
        <th class="text-left">
          大小
        </th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="list in lists"
        :key="list.Name"
      >
        <td>{{ list.Name }}</td>
        <td>{{ list.Size }}</td>
        <v-btn :href="list.Url">下载</v-btn>
      </tr>
    </tbody>
  </v-table>
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
import axios from 'axios';
import { ref } from 'vue';
import { useRoute } from 'vue-router';
const router = useRoute();
const items = [
        {
          title: '下载列表',
          disabled: false,
          href: '/',
        },
        {
          title: router.params.name,
          disabled: false,
          href: router.params.name,
        },
      ]
    const server = ref(localStorage.getItem('server') || config.defaultServer)
  const lists = ref([])
  const downlists = ref([])
  const imglists = ref([])
  axios({
    method:'get',
    url:server.value + 'list/' + router.params.name
  }).then((res)=>{
    console.log(res.data)
    lists.value = res.data.files
    console.log(lists.value)  
})
  </script>
  