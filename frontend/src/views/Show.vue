<template>
    <div v-if="!errorMessage">
        <h2>Here's the secret:</h2>
        <pre>{{ secret }}</pre>
        <p>Can be viewed {{ viewsLeft }} more time<template v-if="viewsLeft !== 1">s</template></p>
    </div>

    <div v-else>
        <h2>Sorry, that secret could not be found. 😢</h2>
    </div>
</template>

<script>
import axios from 'axios'
export default {
    data() {
        return {
            secret: "",
            viewsLeft: 0,
            errorMessage: "",
        }
    },
    mounted() {
        axios.get("/api/secret?id="+this.$route.params.id).then((r) => {
            this.secret = r.data.Value;
            this.viewsLeft = r.data.RemainingViews;
        }).catch((e) => {
            this.errorMessage = e.response.data.error;
        });
    }
}
</script>
