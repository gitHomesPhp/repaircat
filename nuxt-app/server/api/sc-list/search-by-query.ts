export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const query = useQuery(event)

    const data =  await $fetch(
        config.API
        + `/sc-list/search?page=${query.page}&q=${query.q}&city=${query.city}`
    )

    return Object.values(data)
})