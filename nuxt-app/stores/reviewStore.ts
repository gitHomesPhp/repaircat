import {defineStore} from "pinia";

export const useReviewStore = defineStore(
    'reviewStore',
    {
        state: () => ({
            reviews: []
        }),
        actions: {
            async getReviews(scId: number) {
                this.reviews = await $fetch(`/api/sc/review?scId=${scId}`)
            }
        }
    }
)
