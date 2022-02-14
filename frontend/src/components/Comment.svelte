<script lang="ts">
    import { Comment as CommentType } from "../interfaces/comment";
    import Avatar from './atoms/Avatar.svelte';

    export let comment: CommentType;
</script>

<div class="flex flex-row my-5">
    <div class="basis-1/6">
        <Avatar name={comment.author.name}/>
    </div>
    <div class="flex flex-col mx-2">
        <div class="text-lg font-bold"> { comment.author.name } </div>
        <div class="text-base">
            { comment.body }
        </div>
    </div>
</div>