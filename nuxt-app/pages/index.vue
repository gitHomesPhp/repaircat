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
  const page = ref(1)
  const incrementPage = () => page.value++
  const decrementPage = () => page.value--

  const previous = ref(false)
  const next = ref(true)

  const setPaginateNextPrevious = (data) => {
    previous.value = !!data.previous
    next.value = !!data.next
    console.log(next.value)
  }

</script>

<style scoped lang="scss">

</style>