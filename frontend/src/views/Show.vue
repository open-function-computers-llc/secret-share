<template>
    <div v-if="!errorMessage">
        <h2>Here's the secret:</h2>
        <pre>{{ secret }}</pre>
        <div class="copy-row">
            <a class="button button--copy" :class="{ copied: copied }" href="#" @click.prevent="copySecret"><span class="sizer" aria-hidden="true">Copy to clipboard</span><span class="label">{{ copied ? 'Copied!' : 'Copy to clipboard' }}</span></a>
        </div>
        <p>Can be viewed {{ viewsLeft }} more time<template v-if="viewsLeft !== 1">s</template></p>
        <p>Can be viewed until {{ timeLeft }}<template v-if="timeLeft > Date.now"></template></p>
    </div>

    <div v-else>
        <h2>{{ expired ? 'Sorry, that secret has expired!' : 'Sorry, that secret could not be found. 😢' }}</h2>
    </div>
</template>

<script>
import axios from 'axios'
export default {
    data() {
        return {
            secret: "",
            viewsLeft: 0,
            timeLeft: 0,
            errorMessage: "",
            expired: false,
            copied: false,
        }
    },
    methods: {
        copySecret() {
            navigator.clipboard.writeText(this.secret);
            this.copied = true;
            clearTimeout(this.copyTimeout);
            this.copyTimeout = setTimeout(() => {
                this.copied = false;
            }, 2000);
        }
    },
    mounted() {
        axios.get("/api/secret?id="+this.$route.params.id).then((r) => {
            this.secret = r.data.Value;
            this.viewsLeft = r.data.RemainingViews;
            this.timeLeft = r.data.TimeString;
            
        }).catch((e) => {
            this.errorMessage = e.response?.data?.error || "Secret not found";
            this.expired = e.response?.data?.status === "expired";
        });
    }
}
</script>

<style scoped>
.copy-row {
    margin: 0 0 1.25rem;
}
</style>
