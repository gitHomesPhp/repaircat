export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const query = useQuery(event)

    const data =  await $fetch(config.API + '/sc-list?page=' + query.page)

    return Object.values(data)
})