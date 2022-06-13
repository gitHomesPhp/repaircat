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
      :address="sc.address"
      :underground="sc.underground"
    />
  </div>
</template>

<script lang="ts" setup>
const props = defineProps({
  page: { type: Number, required: true }
})

const emit = defineEmits([
  'paginate-next-previous'
])

const fetchList = async (value) => {
  let scList = await $fetch('/api/sc-list?page=' + value)
  let paginate = scList.pop()
  emit("paginate-next-previous", paginate)
  return scList
}

const data = await fetchList(props.page)


const scList = ref(data)

watch(() => props.page, async (value) => {
  scList.value = await fetchList(value)
})

</script>

<style scoped>

</style>