env = import_module("env.star")
deployer = import_module("deployer.star")
types = import_module("types.star")

geth = import_module("components/geth.star")
batcher = import_module("components/batcher.star")
node = import_module("components/node.star")
proposer = import_module("components/proposer.star")

# TODO: Make these configurable
GETH_RPC_URL="http://localhost:8545"
NODE_RPC_URL="http://localhost:8547"
BATCHER_RPC_URL="http://localhost:8548"
PROPOSER_RPC_URL="http://localhost:8560"

def launch_op(plan, images, l1_rpc_url):
    l1 = types.get_l1(l1_rpc_url)
    l2 = types.get_l2(GETH_RPC_URL)
    files = types.get_files(plan)
    e = env.get(plan, files, l1)

    deployer.fund_wallets(plan, e)
    deployer.build_deploy_config(plan, e, l1.chain_id)
    # deploy create2 factory

    deployer.deploy_l1_contracts(plan, e, files)
    node.init(plan, images["node"], e)
    deployer.generate_jwt_secret(plan, e)
    geth.init(plan, image["geth"], files)

    # Deploy L2 Components
    geth.launch(plan, image["geth"], l1, files)
    node.launch(plan, image["node"], files, e, l1, l2, NODE_RPC_URL)
    batcher.launch(plan, image["batcher"], e, l1, l2, BATCHER_RPC_URL, NODE_RPC_URL)
    proposer.launch(plan, image["proposer"], e, l1, PROPOSER_RPC_URL, NODE_RPC_URL)

    # TODO: Bridge Tokens to Address