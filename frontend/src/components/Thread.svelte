<script lang="ts">
    import { Comment as CommentType } from "../interfaces/comment";
    import Comment from './Comment.svelte';

    export let comments = [];
</script>

<div class="flex flex-col bg-white shadow-md border rounded-lg p-5">
    <div>
	    {#each comments as comment}
            <Comment comment={comment}/>
	    {/each}
    </div>
</div>