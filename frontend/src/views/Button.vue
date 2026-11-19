<template>
    <div v-if="!errorMessage">
       <a class = "button" :href = "'/#/show=true/' + secretID">Click here to see your secret</a>
    </div>

    <div v-else>
        <h2>Sorry, that secret could not be found. 😢</h2>
    </div>
</template>

<script>
export default {
    data() {
        return {
            secretID: this.$route.params.id,
            errorMessage: "",
        }
    },
}
</script>
<style lang="css" scoped>
.button {
  margin: 48px auto 0;
  width: min(380px, 80%);
}
</style>