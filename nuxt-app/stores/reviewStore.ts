import {defineStore} from "pinia";

export const useReviewStore = defineStore(
    'reviewStore',
    {
        state: () => ({
            reviews: []
        }),
        actions: {
            async getReviews(page: number, scId: number) {
                const data = await $fetch(`/api/sc/review?scId=${scId}&page=${page}`)
                this.reviews = this.reviews.concat(data)
            }
        }
    }
)
