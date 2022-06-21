export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    console.log(config.API)
    console.log(config.API + '/city')
    const data = await $fetch(config.API + '/city')

    return data
})