import subprocess
import argparse
import json
import re
import subprocess
import sys

from pycoin import encoding
from pycoin.ecdsa.secp256k1 import secp256k1_generator
from pycoin.serialize import b2h, h2b
from pycoin.key import Key
from pycoin.key.key_from_text import key_from_text
from pycoin.key.BIP32Node import BIP32Node
from pycoin.networks import full_network_name_for_netcode, network_name_for_netcode, network_codes
from pycoin.networks.default import get_current_netcode
from pycoin.ui import address_for_pay_to_script
from pycoin.tx.pay_to.ScriptPayToAddressWit import ScriptPayToAddressWit


def gpg_entropy():
    try:
        output = subprocess.Popen(
            ["gpg", "--gen-random", "2", "64"], stdout=subprocess.PIPE).communicate()[0]
        return output
    except OSError:
        sys.stderr.write("warning: can't open gpg, can't use as entropy source\n")
    return b''


def get_entropy():
    entropy = bytearray()
    try:
        entropy.extend(gpg_entropy())
    except Exception:
        print("warning: can't use gpg as entropy source", file=sys.stdout)
    try:
        entropy.extend(open("/dev/random", "rb").read(64))
    except Exception:
        print("warning: can't use /dev/random as entropy source", file=sys.stdout)
    entropy = bytes(entropy)
    if len(entropy) < 64:
        raise OSError("can't find sources of entropy")
    return entropy

def _create_bip32(_):
    max_retries = 64
    for _ in range(max_retries):
        try:
            return BIP32Node.from_master_secret(get_entropy(), netcode='BTC')
        except ValueError as e:
            continue
    # Probably a bug if we get here
    raise RuntimeError("can't create BIP32 key")


k2 = Key.from_text(_create_bip32)
print(k2)
