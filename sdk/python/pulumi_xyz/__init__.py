# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .random import *
_utilities.register(
    resource_modules="""
[
 {
  "pkg": "render",
  "mod": "index",
  "fqn": "pulumi_render",
  "classes": {
   "render:index:Random": "Random"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "render",
  "token": "pulumi:providers:render",
  "fqn": "pulumi_render",
  "class": "Provider"
 }
]
"""
)
