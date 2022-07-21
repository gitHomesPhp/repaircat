<template>
  <div class="container">
    <AppHead/>
    <AppBody>
      <CommonFinder />
      <Paginator v-if="scList.length" type="top"/>
      <ScList />
      <Paginator v-if="scList.length" type="bottom"/>
    </AppBody>
    <AppFoot/>
  </div>
</template>

<script setup lang="ts">
definePageMeta({
  middleware: "paginate-after-find-sc-by-underground"
})


import {useScPaginatorStore} from "~/stores/scPaginatorStore";
const { scList } = useScPaginatorStore()
</script>

<style scoped lang="scss">

</style>