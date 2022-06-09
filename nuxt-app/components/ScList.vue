<template>
  <div>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
    <ScCard :name="sc.Name" :description="sc.Description" :phone="sc.Phone" :email="sc.Email" :site="sc.Site"/>
  </div>
</template>

<script lang="ts" setup>
const scList = await $fetch('/api/sc-list')

const sc = scList[0];
</script>

<style scoped>

</style>