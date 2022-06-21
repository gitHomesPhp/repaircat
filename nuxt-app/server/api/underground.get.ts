export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()
    const query = useQuery(event)

    const data = await $fetch(config.API + `/city/${query.city}/underground`)

    return data
})