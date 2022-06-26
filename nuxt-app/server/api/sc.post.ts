export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const body = await useBody(event)

    const formData = new FormData()

    for ( let key in body ) {
        if (key === 'undergrounds') {
            body[key].forEach((underground, index, array) => {
                let undKey = `undergrounds`
                formData.append(undKey, underground.id);
                console.log(undKey, underground.id)
            })

            continue
        }
        formData.append(key, body[key]);
    }

    const data =  await $fetch(config.API + '/sc', {
        method: 'post',
        body: formData,
    })

    return Object.values(data)
})