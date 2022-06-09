<template>
  <main>
    <div class="content">
      <slot />
    </div>
  </main>
</template>

<script lang="ts" setup>


</script>

<style scoped>
  main {
    flex-grow: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
  }
  .content {
    width: 80vw;
  }
</style>