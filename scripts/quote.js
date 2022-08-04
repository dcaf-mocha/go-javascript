const {Keypair, Connection, PublicKey} = require("@solana/web3.js");
const {AnchorProvider} = require("@project-serum/anchor");
const {AccountFetcher, WhirlpoolContext, ORCA_WHIRLPOOL_PROGRAM_ID, PDAUtil, buildWhirlpoolClient,
    swapQuoteByInputToken
} = require("@orca-so/whirlpools-sdk");
const NodeWallet = require("@project-serum/anchor/dist/cjs/nodewallet");
const {u64} = require("@solana/spl-token");
const {Percentage} = require("@orca-so/common-sdk");

const args = process.argv.slice(2);
const config = new PublicKey(args[0]);
const tokenAMint = new PublicKey(args[1]);
const tokenBMint = new PublicKey(args[2]);
const inputToken = new PublicKey(args[3]);

async function getQuote() {

    const wallet = new NodeWallet.constructor(Keypair.generate());
    const provider = new AnchorProvider(
        new Connection("https://api.devnet.solana.com", "confirmed"),
        wallet,
        AnchorProvider.defaultOptions()
    );

    const fetcher = new AccountFetcher(provider.connection);
    const ctx = WhirlpoolContext.withProvider(provider, ORCA_WHIRLPOOL_PROGRAM_ID);

    const whirlpoolPda = PDAUtil.getWhirlpool(
        ORCA_WHIRLPOOL_PROGRAM_ID,
        config,
        tokenAMint,
        tokenBMint,
        // new PublicKey("CRR7huZnXaiBjGGMAU6iVeQU9b2g71NXiLHA6g29DeYN"),
        // new PublicKey("57K3gMtUMctYGYUpm9PjzYQeiCV8BeRkSuuBFGkuWAdt"),
        // new PublicKey("Dphoc5nPvC5eadUP79McRB36hgKcetgJ7BRG5Zv6QeYp"),
        64);

    const whirlpoolClient = buildWhirlpoolClient(ctx, fetcher);
    const whirlpool = await whirlpoolClient.getPool(whirlpoolPda.publicKey, true);
    // const whirlpoolData = await whirlpool.getData();

    const swapQuote =  await swapQuoteByInputToken(
        whirlpool,
        inputToken,
        new u64(100),
        Percentage.fromFraction(0, 100),
        ORCA_WHIRLPOOL_PROGRAM_ID,
        fetcher,
        true,
    );
    const swapQuoteString =  {
        estimatedAmountIn: swapQuote.estimatedAmountIn.toString(),
        estimatedAmountOut: swapQuote.estimatedAmountOut.toString(),
        estimatedEndTickIndex: swapQuote.estimatedEndTickIndex,
        estimatedEndSqrtPrice: swapQuote.estimatedEndSqrtPrice.toString(),
        estimatedFeeAmount: swapQuote.estimatedFeeAmount.toString(),
        amount: swapQuote.amount.toString(),
        amountSpecifiedIsInput: swapQuote.amountSpecifiedIsInput,
        aToB: swapQuote.aToB,
        otherAmountThreshold: swapQuote.otherAmountThreshold.toString(),
        sqrtPriceLimit: swapQuote.sqrtPriceLimit.toString(),
        tickArray0: swapQuote.tickArray0.toString(),
        tickArray1: swapQuote.tickArray1.toString(),
        tickArray2: swapQuote.tickArray2.toString(),
    };
    console.log(JSON.stringify(swapQuoteString));
    return JSON.stringify(swapQuoteString);
}
return getQuote();
