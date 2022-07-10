export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const data = await $fetch(config.API + '/cities')

    return data
})