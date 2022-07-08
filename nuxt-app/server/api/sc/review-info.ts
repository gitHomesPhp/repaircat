export default defineEventHandler(async (event) => {
    const config = useRuntimeConfig()

    return await $fetch(config.API + `/sc/${event.context.params.id}/review-info`)
})