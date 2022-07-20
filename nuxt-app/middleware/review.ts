import { useReviewStore } from "~/stores/reviewStore";

export default defineNuxtRouteMiddleware(async (to, from) => {
    const { getReviews } = useReviewStore()

    if (process.client) {
        await getReviews(Number(to.params.id));
    }
})