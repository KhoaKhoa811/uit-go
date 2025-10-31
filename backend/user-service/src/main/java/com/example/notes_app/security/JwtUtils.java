package com.example.notes_app.security;

import java.util.Base64;
import java.util.Date;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.security.core.Authentication;
import org.springframework.stereotype.Component;

import com.example.notes_app.enums.ErrorCode;
import com.example.notes_app.exception.impl.InvalidTokenSignatureException;
import com.nimbusds.jose.JOSEException;
import com.nimbusds.jose.JWSAlgorithm;
import com.nimbusds.jose.JWSHeader;
import com.nimbusds.jose.JWSObject;
import com.nimbusds.jose.JWSVerifier;
import com.nimbusds.jose.Payload;
import com.nimbusds.jose.crypto.MACSigner;
import com.nimbusds.jose.crypto.MACVerifier;
import com.nimbusds.jwt.JWTClaimsSet;
import com.nimbusds.jwt.SignedJWT;

import lombok.RequiredArgsConstructor;

@Component
@RequiredArgsConstructor
public class JwtUtils {

    @Value("${jwt.signed-key}")
    private String SIGNER_KEY;

    // Access token: 1 hour
    private static final long ACCESS_TOKEN_EXPIRATION = 3600000; // 1h = 60 * 60 * 1000
    // Refresh token: 30 days
    private static final long REFRESH_TOKEN_EXPIRATION = 2592000000L; // 30 * 24 * 60 * 60 * 1000

    /**
     * Generate Access Token (no role, only authentication)
     */
    public String generateAccessToken(Authentication authentication) throws JOSEException {
        String email = authentication.getName();

        JWTClaimsSet claims = new JWTClaimsSet.Builder()
                .subject(email)
                .issueTime(new Date())
                .expirationTime(new Date(System.currentTimeMillis() + ACCESS_TOKEN_EXPIRATION))
                .build();

        return signJwt(claims);
    }

    /**
     * Generate Refresh Token (no role, only authentication)
     */
    public String generateRefreshToken(Authentication authentication) throws JOSEException {
        String email = authentication.getName();

        JWTClaimsSet claims = new JWTClaimsSet.Builder()
                .subject(email)
                .issueTime(new Date())
                .expirationTime(new Date(System.currentTimeMillis() + REFRESH_TOKEN_EXPIRATION))
                .build();

        return signJwt(claims);
    }

    /**
     * Extract email (subject) from JWT
     */
    public String getEmail(String token) {
        try {
            if (isValidToken(token)) {
                SignedJWT signedJWT = SignedJWT.parse(token);
                return signedJWT.getJWTClaimsSet().getSubject();
            } else {
                throw new InvalidTokenSignatureException(ErrorCode.JWT_INVALID);
            }
        } catch (Exception e) {
            throw new RuntimeException("Error while validating JWT", e);
        }
    }

    /**
     * Validate JWT signature and expiration
     */
    public boolean isValidToken(String token) {
        try {
            SignedJWT signedJWT = SignedJWT.parse(token);
            byte[] decodedKey = Base64.getDecoder().decode(SIGNER_KEY);
            JWSVerifier verifier = new MACVerifier(decodedKey);

            boolean signatureValid = signedJWT.verify(verifier);
            boolean notExpired = signedJWT.getJWTClaimsSet().getExpirationTime().after(new Date());

            return signatureValid && notExpired;
        } catch (Exception e) {
            return false;
        }
    }

    /**
     * Helper method for signing a JWT
     */
    private String signJwt(JWTClaimsSet claims) throws JOSEException {
        JWSHeader header = new JWSHeader(JWSAlgorithm.HS256);
        JWSObject jwsObject = new JWSObject(header, new Payload(claims.toJSONObject()));

        byte[] decodedKey = Base64.getDecoder().decode(SIGNER_KEY);
        jwsObject.sign(new MACSigner(decodedKey));

        return jwsObject.serialize();
    }
}
