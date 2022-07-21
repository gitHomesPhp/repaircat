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
            async fetchScListByUnderground(page, city, underground) {
                const scList = await $fetch(
                    `/api/sc-list/search-by-underground?page=${page}&underground=${underground}&city=${city}`
                )
                this.paginate = scList.pop()
                this.scList = scList.pop()
            },
            async fetchScListByMunicipality(page, city, municipality) {
                const scList = await $fetch(
                    `/api/sc-list/search-by-municipality?page=${page}&municipality=${municipality}&city=${city}`
                )
                this.paginate = scList.pop()
                this.scList = scList.pop()
            },
            async fetchScListByQuery(page, city, query) {
                const scList = await $fetch(
                    `/api/sc-list/search-by-query?page=${page}&q=${query}&city=${city}`
                )
                this.paginate = scList.pop()
                this.scList = scList.pop()
            }
        },
        getters: {
            nextPage: (state) => state.currentPage + 1,
            previousPage: (state) => state.currentPage - 1,
            canNextPage: (state) => state.paginate.next,
            canPreviousPage: (state) => state.paginate.previous,
        }
    })
