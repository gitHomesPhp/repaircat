export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()

    const data =  await $fetch(config.API + '/sc-list?page=1')

    return data
})