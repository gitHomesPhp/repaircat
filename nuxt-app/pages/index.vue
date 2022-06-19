<template>
  <div class="container">
    <AppHead />
    <AppBody>
      <Paginator
          @next-page="incrementPage"
          @previous-page="decrementPage"
          type="top"
          :page="page"
          :next="next"
          :previous="previous"
      />
      <ScList @paginate-next-previous="setPaginateNextPrevious" :page="page" />
      <Paginator
          @next-page="incrementPage"
          @previous-page="decrementPage"
          type="bottom"
          :page="page"
          :next="next"
          :previous="previous"
      />
    </AppBody>
    <AppFoot />
  </div>
</template>

<script lang="ts" setup>
  const route = useRoute()
  const router = useRouter()
  const page = ref(route.query.page ? route.query.page: 1)


  const incrementPage = () => {
    page.value++
    router.push({
      query: {
        page: page.value
      }
    })
  }
  const decrementPage = () => {
    page.value--
    router.push({
      query: {
        page: page.value
      }
    })
  }

  const previous = ref(false)
  const next = ref(true)

  const setPaginateNextPrevious = (data) => {
    previous.value = !!data.previous
    next.value = !!data.next
  }

  watch(() => route.query.page, () => page.value = route.query.page ? route.query.page: 1)

</script>

<style scoped lang="scss">

</style>