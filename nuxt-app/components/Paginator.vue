<template>
  <div class="paginator" :class="type + '-paginator'">
    <div class="paginator__wrapper">
      <span>
        <NuxtLink v-if="canPreviousPage"
                  :to="`/?page=${previousPage}`"
                  class="pointer"
        >
          Предыдущая
        </NuxtLink>
      </span>
      <span class="paginator__page">{{ currentPage }}</span>
      <span>
        <NuxtLink v-if="canNextPage"
                  :to="`/?page=${nextPage}`"
                  class="pointer"
        >
          Следующая
        </NuxtLink>
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
  import {useScPaginatorStore} from "~/stores/scPaginatorStore";
  import {storeToRefs} from "pinia";

  const { canPreviousPage, canNextPage, nextPage, previousPage, currentPage } = storeToRefs(useScPaginatorStore())

  const props = defineProps({
    type: { type: String },
  })

</script>

<style scoped lang="scss">
a {
  text-decoration: none;
  color: #0070c0;
}
.paginator {
  display: flex;
  &__wrapper {
    width: 100%;
    display: flex;
    justify-content: space-between;
  }
  &__page {
    font-weight: bolder;
  }
}
.top-paginator {
  margin-bottom: 1rem;
  padding-bottom: 1rem;
  border-bottom: #cccccc 1px solid;
}
.bottom-paginator {
  margin-top: 1rem;
  padding-top: 1rem;
  border-top: #cccccc 1px solid;
}
</style>