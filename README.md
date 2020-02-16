[![](https://jitpack.io/v/dev.reimer/dex-api.svg)](https://jitpack.io/#dev.reimer/dex-api)

# 📇 dex-api<sup>[α](#status-α)</sup>

Java wrapper for the [Dex](https://github.com/dexidp/dex) OIDC provider [gRPC](https://grpc.io/) API.

## Gradle Dependency

This library is available on [**jitpack.io**](https://jitpack.io/#dev.reimer/dex-api).  
Add this in your `build.gradle.kts` or `build.gradle` file:

<details open><summary>Kotlin</summary>

```kotlin
repositories {
    maven("https://jitpack.io")
}

dependencies {
    implementation("dev.reimer:dex-api:<version>")
}
```

</details>

<details><summary>Groovy</summary>

```groovy
repositories {
    maven { url 'https://jitpack.io' }
}

dependencies {
    implementation 'dev.reimer:dex-api:<version>'
}
```

</details>

## Status α

⚠️ _Warning:_ This project is in an experimental alpha stage:
- The API may be changed at any time without further notice.
- Development still happens on `master`.
- Pull Requests are highly appreciated!