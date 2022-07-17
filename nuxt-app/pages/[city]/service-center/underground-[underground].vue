<template>
  <div class="container">
    <AppHead/>
    <AppBody>
      <CommonFinder />
      <Paginator type="top"/>
      <ScList />
      <Paginator type="bottom"/>
    </AppBody>
    <AppFoot/>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: "paginate-after-find-sc-by-underground"
})

const route = useRoute()

</script>

<style scoped lang="scss">

</style>