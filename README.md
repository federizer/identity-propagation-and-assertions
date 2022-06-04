<!-- @import "style.less" -->

# Identity Propagation and Assertions

## Introduction

With the growing popularity of protocols based on the OAuth 2.0 specification, there is a need for an interoperable standard that specifies how to convey information about the user from an identity provider (IdP) to a resource server (RS) across security domain boundaries.

## Motivation

To allow an IdP to provide authentication attributes to a number of separately-administered RSs as illustrated in Figure&nbsp;1

![Model](./images/identity_propagation_model.svg)

<p class="figure">
Fig.&nbsp;1.&emsp;Identity Propagation Model
</p>

## Sequence Diagram

The sequence diagram illustrated in Figure&nbsp;2 shows an identity propagation flow for the subscriber authenticated to the IdP to be able to access the resources stored on the RS. The following are prerequisites for the given scenario:

TBD

<div class="diagram">
    <img src=./images/identity_propagation_flow.svg alt="Sequence Diagram">
</div>

<p class="figure">
Fig.&nbsp;2.&emsp;Identity Propagation Flow
</p>

## Usability Considerations

Using OAuth 2.0 and OpenID Connect technologies is an effective option to secure web APIs. From an API ecosystem perspective, the resource server is the API provider, and the client is the API consumer. The identity propagation system is helpful in B2B scenarios.

#### Trust First

The primary benefit of the identity propagation system is to address the problem of trust between the identity provider, the API provider, and the API consumer.

#### Third-Party Users to API

To identify the third-party users and clients, use composite API tokens generated by the identity propagation system instead of the API keys, user IDs, and passwords.

#### User to Third-Party APIs

Use a composite API token identifying the user and the client itself generated by the identity propagation system to access third-party APIs. Do not use an API key, user ID, and password.
