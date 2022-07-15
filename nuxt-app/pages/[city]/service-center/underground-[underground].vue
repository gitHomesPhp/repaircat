<template>
  <div class="container">
    <AppHead/>
    <AppBody>
      {{route.params}}
    </AppBody>
    <AppFoot/>
  </div>
</template>

<script setup lang="ts">
  const route = useRoute()

</script>

<style scoped lang="scss">

</style>