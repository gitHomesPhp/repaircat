<template>
  <div>
    <ScCard
      v-for="sc in scList"
      :key="sc.id"
      :name="sc.name"
      :description="sc.description"
      :phone="sc.phone"
      :email="sc.email"
      :site="sc.site"
    />
  </div>
</template>

<script lang="ts" setup>
const props = defineProps({
  page: { type: Number, required: true }
})

const data = await $fetch('/api/sc-list?page=' + props.page)

const scList = ref(data)

watch(() => props.page, async (value) => {
  scList.value = await $fetch('/api/sc-list?page=' + value)
})

</script>

<style scoped>

</style>