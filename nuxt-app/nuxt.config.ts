import { defineNuxtConfig } from 'nuxt'

export default defineNuxtConfig({
    css: [
        'assets/css/reset.css'
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
})
