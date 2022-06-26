<template>
  <div v-if="show" class="submenu">
    <slot />
  </div>
</template>

<script setup lang="ts">
  const props = defineProps({
    show: { type: Boolean, default: false }
  })
</script>

<style scoped lang="scss">
  .submenu {
    position: absolute;
    top: 1rem;
    border: 0 solid rgba(0,0,0,.125);
    box-shadow: 0 0 1px rgb(0 0 0 / 13%), 0 1px 3px rgb(0 0 0 / 20%);
    background-color: rgba(255,255,255,.9);
    z-index: 1;
    width: 130px;
    border-radius: 5px;
  }
</style>