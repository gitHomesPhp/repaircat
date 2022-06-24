<template>
  <div style="height: 100vh">
    <NuxtPage />
  </div>
</template>

<script setup lang="ts">

useHead({
  title: 'Ремонт техники',
  viewport: 'width=device-width, initial-scale=1, maximum-scale=1',
  charset: 'utf-8',
  description: 'Ремонт техники',
  htmlAttrs: {
    lang: 'ru'
  },
  link: [
    { rel: 'apple-touch-icon', sizes: '180x180', href: '/apple-touch-icon.png' },
    { rel: 'icon', type: 'image/png', sizes: '32x32', href: '/favicon-32x32.png' },
    { rel: 'icon', type: 'image/png', sizes: '16x16', href: '/favicon-16x16.png' },
    { rel: 'manifest', href: '/site.webmanifest' },
    { rel: 'mask-icon', href: '/safari-pinned-tab.svg', color: '#5bbad5'},
  ],
  meta: [
    { name: 'msapplication-TileColor', content: '#da532c' },
    { name: 'theme-color', content: '#ffffff' },
  ]
})

</script>

<style>
.container {
  display: flex;
  flex-direction: column;
  min-height: 100%;
}
.ymap-container {
  height: 100%;
}
</style>