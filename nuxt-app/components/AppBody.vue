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
    padding: 0 1rem;
    background-color: #fafafa;
  }
  .content {
    width: 100%;
    max-width: 870px;
    padding-top: 1rem;
    padding-bottom: 1rem;

    flex-grow: 1;
  }
</style>