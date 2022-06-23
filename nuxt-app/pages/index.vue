<template>
  <div class="container">
    <AppHead />
    <AppBody>
      <Paginator type="top"/>
      <ScList />
      <Paginator type="bottom"/>
    </AppBody>
    <AppFoot />
  </div>
</template>

<script lang="ts" setup>
  definePageMeta({
    middleware: "paginate-middleware"
  })
</script>

<style scoped lang="scss">

</style>