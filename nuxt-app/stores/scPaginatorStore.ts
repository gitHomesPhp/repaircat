import { defineStore } from "pinia";

export const useScPaginatorStore = defineStore(
    'scPaginatorStore',
    {
        state: () => ({
            currentPage: 1,
            scList: [],
            paginate: {},
        }),
        actions: {
            setPage(page) {
                this.currentPage = page
            },
            async fetchScList(value) {
                const scList = await $fetch('/api/sc-list?page=' + value)
                this.paginate = scList.pop()
                this.scList = scList.pop()
            },
        },
        getters: {
            nextPage: (state) => state.currentPage + 1,
            previousPage: (state) => state.currentPage - 1,
            canNextPage: (state) => state.paginate.next,
            canPreviousPage: (state) => state.paginate.previous,
        }
    })
