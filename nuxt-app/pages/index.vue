<template>
  <div class="container">
    <AppHead />
    <AppBody>
      <Paginator
          @next-page="incrementPage"
          @previous-page="decrementPage"
          type="top"
          :page="page"
      />
      <ScList :page="page" />
      <Paginator
          @next-page="incrementPage"
          @previous-page="decrementPage"
          type="bottom"
          :page="page"
      />
    </AppBody>
    <AppFoot />
  </div>
</template>

<script lang="ts" setup>
  const page = ref(1)
  const incrementPage = () => page.value++
  const decrementPage = () => page.value--

</script>

<style scoped lang="scss">

</style>