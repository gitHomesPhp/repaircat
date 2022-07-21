<template>
  <div class="container">
    <AppHead />
    <AppBody>
      <div>
        <CommonFinder />
        <!--Paginator type="top"/>
        <ScList />
        <Paginator type="bottom"/-->
      </div>
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