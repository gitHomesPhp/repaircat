<template>
  <div class="container">
    <AppHead/>
    <AppBody>
      <Sc />
    </AppBody>
    <AppFoot/>
  </div>
</template>

<script lang="ts" setup>
definePageMeta({
  middleware: 'review'
})
const route = useRoute();

</script>

<style scoped lang="scss">
</style>