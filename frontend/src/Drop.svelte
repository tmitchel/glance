<script>
    import { onMount } from 'svelte';
    onMount(async () => {
        window.$('.ui.dropdown').dropdown({showOnFocus: false});
    })
    export let statuses;
    export let card;
</script>

<div class="ui selection dropdown">
    <input type="hidden" name="gender">
    <i class="dropdown icon"></i>
    <div class="default text">{statuses[card.status]}</div>
    <div class="menu">
    {#each statuses as status, i}
        <div class="item" data-value={i}>{status}</div>
    {/each}
    </div>
</div>