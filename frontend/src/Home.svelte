<script>
    let hey = "hey";

    import { CardService } from "../../generated/oto.gen.js";
    const cardService = new CardService();
    let meh = {};
	cardService.cards({}).then(resp => {
        meh = resp;
    }).catch(err => console.log(err))

</script>

<h1>{hey}</h1>
<h2>mey</h2>
<pre>{meh}</pre>
<h1>you</h1>