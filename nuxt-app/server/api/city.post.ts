export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const body = await useBody(event)

    const formData = new FormData()

    for ( let key in body ) {
        formData.append(key, body[key]);
    }

    const data =  await $fetch(config.API + '/city', {
        method: 'post',
        body: formData,
    })

    return Object.values(data)
})