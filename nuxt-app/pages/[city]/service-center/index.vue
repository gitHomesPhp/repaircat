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

<script lang="ts" setup>
definePageMeta({
  middleware: 'paginate-after-find-by-query',
})


import {useScPaginatorStore} from "~/stores/scPaginatorStore";
const { scList } = useScPaginatorStore()
</script>

<style scoped>

</style>