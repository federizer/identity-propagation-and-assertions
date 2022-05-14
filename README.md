<!-- @import "style.less" -->

# Identity Propagation

## Introduction

&emsp;With the growing popularity of protocols based on the OAuth 2.0 specification, there is a need for an interoperable standard that specifies how to convey information about the user from an identity provider (IdP) to a relying party (RP) across security domain boundaries.

## Motivation

&emsp;To allow an IdP to provide authentication attributes to a number of separately-administered RPs as illustrated in Figure&nbsp;1

![Model](./images/identity_propagation_model.svg)

<p class="figure">
Fig.&nbsp;1.&emsp;Identity Propagation Model
</p>

## Sequence Diagram

&emsp;The sequence diagram illustrated in Figure&nbsp;2 shows an identity propagation flow for the subscriber authenticated to the IdP to be able to access the resources stored on the RP. The following are prerequisites for the given scenario:

TBD

![Sequence Diagram](./images/identity_propagation_flow.svg)

<p class="figure">
Fig.&nbsp;2.&emsp;Identity Propagation Flow
</p>


