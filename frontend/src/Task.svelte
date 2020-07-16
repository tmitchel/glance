<script>
    import { onMount } from 'svelte';
    onMount(async () => {
        window.$('.ui.dropdown').dropdown({showOnFocus: false});
    })

    export let card;
    export let volunteer;
    export let createService;

    let status = 0;

    let updateStatus = async function(card) {
        createService.updateStatus({
            status: status,
            card: card,
            user: ""
        })
    }

    const statuses = [
        "Not Started",
        "Completed",
        "Help Requested",
        "In Progress",
    ];
</script>

{#if card}
    <div id="task" class="ui modal">
        <div class="content">
            <h1 class="header">{card.title || ""}</h1>
            <p>{card.content || ""}</p>
        </div>
        <div class="actions">
            {#if card.claimed}
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
                <button button class="ui positive button" type="submit" on:click={updateStatus(card)}>Submit</button>
            {:else}
                <button class="ui positive button" type="submit" on:click={volunteer("", card.id)}>Volunteer</button>
            {/if}
            <button class="ui negative button">Cancel</button>
        </div>
    </div>
{/if}
