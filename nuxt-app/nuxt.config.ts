import { defineNuxtConfig } from 'nuxt'

export default defineNuxtConfig({
    css: [
        'assets/css/reset.css',
        'assets/css/global.css'
    ],
    scss: {
        implementation: require('sass'),
    },
    privateRuntimeConfig: {
        API: process.env.API_BASE_URL
    },
    publicRuntimeConfig: {
        API: process.env.API_BASE_URL
    },
    buildModules: ['@pinia/nuxt'],
    plugins: [
        { src: '~/plugins/ymapPlugin.ts',  mode: 'client' }
    ]
})
