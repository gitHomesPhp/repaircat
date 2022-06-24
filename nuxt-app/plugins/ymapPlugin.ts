import YmapPlugin from 'vue-yandex-maps'

export default defineNuxtPlugin((nuxtApp) => {
     nuxtApp.vueApp.use(YmapPlugin, {
        apiKey: 'a4023680-dd43-4ac2-b4da-f2eeebf38db9',
        lang: 'ru_RU',
        coordorder: 'latlong',
        enterprise: false,
        version: '2.1'
    });
});