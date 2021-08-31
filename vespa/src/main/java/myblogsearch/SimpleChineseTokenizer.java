package myblogsearch;

import com.huaban.analysis.jieba.SegToken;
import com.huaban.analysis.jieba.JiebaSegmenter;
import com.huaban.analysis.jieba.JiebaSegmenter.SegMode;
import com.yahoo.language.Language;
import com.yahoo.language.LinguisticsCase;
import com.yahoo.language.process.Normalizer;
import com.yahoo.language.process.StemMode;
import com.yahoo.language.process.Token;
import com.yahoo.language.process.Tokenizer;
import com.yahoo.language.process.TokenType;
import com.yahoo.language.process.Transformer;
import com.yahoo.language.simple.SimpleNormalizer;
import com.yahoo.language.simple.SimpleToken;
import com.yahoo.language.simple.SimpleTokenType;
import com.yahoo.language.simple.SimpleTransformer;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Set;
import opennlp.tools.stemmer.PorterStemmer;
import opennlp.tools.tokenize.SimpleTokenizer;
import opennlp.tools.util.Span;

public class SimpleChineseTokenizer implements Tokenizer {

    private final Set<String> stopwords;
    private final JiebaSegmenter segmenter;
    private final Normalizer normalizer = new SimpleNormalizer();
    private final Transformer transformer = new SimpleTransformer();
    private final PorterStemmer stemmer = new PorterStemmer();

    public SimpleChineseTokenizer(Set<String> stopwords) {
        this.stopwords = stopwords;
        this.segmenter = new JiebaSegmenter();
    }

    @Override
    public Iterable<Token> tokenize(String input, Language language, StemMode stemMode, boolean removeAccents) {
        if (input.isEmpty()) return List.of();

        List<Token> tokens = new ArrayList<Token>();
        List<Span> segments = isChinese(language) ? tokenizeChinese(input) : tokenizeOthers(input);
        for (Span segment : segments) {
            String original = input.substring(segment.getStart(), segment.getEnd());
            String token = normalizer.normalize(original);
            token = LinguisticsCase.toLowerCase(token);
            if (removeAccents) {
                token = transformer.accentDrop(token, language);
            }
            if (stemMode != StemMode.NONE) {
                token = stemmer.stem(token);
            }
            TokenType tokenType = SimpleTokenType.valueOf(input.codePointAt(segment.getStart()));
            tokens.add(new SimpleToken(original).setOffset(segment.getStart())
                                                .setType(tokenType)
                                                .setTokenString(token));
        }
        return tokens;
    }

    private boolean isChinese(Language language) {
        return (language == Language.CHINESE_SIMPLIFIED || language == Language.CHINESE_TRADITIONAL);
    }

    private List<Span> tokenizeChinese(String input) {
        List<Span> tokens = new ArrayList<Span>();
        for (SegToken token : segmenter.process(input, SegMode.SEARCH)) {
            if (stopwords.contains(token.word)) continue;
            tokens.add(new Span(token.startOffset, token.endOffset));
        }
        return tokens;
    }

    private List<Span> tokenizeOthers(String input) {
        return Arrays.asList(SimpleTokenizer.INSTANCE.tokenizePos(input));
    }
}
