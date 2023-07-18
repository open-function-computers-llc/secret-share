<template>
    <div v-if="!errorMessage">
       <a :href = "'/#/show=true/' + secretID">Click here to see your secret</a>
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